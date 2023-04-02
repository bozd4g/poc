//go:build integration
// +build integration

package todo

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"poc-testing/server/internal/db"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

var dbi *sql.DB

func TestMain(m *testing.M) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatal("could not construct pool", err)
	}

	err = pool.Client.Ping()
	if err != nil {
		log.Fatal("could not connect to Docker", err)
	}

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "11",
		Env: []string{
			"POSTGRES_PASSWORD=secret",
			"POSTGRES_USER=user_name",
			"POSTGRES_DB=dbname",
			"listen_addresses = '*'",
		},
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})

	if err != nil {
		log.Fatal("could not start resource", err)
	}

	hostAndPort := resource.GetHostPort("5432/tcp")
	databaseUrl := fmt.Sprintf("postgres://user_name:secret@%s/dbname?sslmode=disable", hostAndPort)

	resource.Expire(120)
	pool.MaxWait = 120 * time.Second

	if err = pool.Retry(func() error {
		dbi, err = sql.Open("postgres", databaseUrl)
		if err != nil {
			return err
		}
		return dbi.Ping()
	}); err != nil {
		log.Fatal("could not connect to docker", err)
	}


	err = db.MigrateUp(dbi)
	if err != nil {
		log.Fatal("failed to run migrations", err)
	}

	code := m.Run()

	if err := pool.Purge(resource); err != nil {
		log.Fatal("Could not purge resource", err)
	}

	os.Exit(code)
}

func Test_Integration_CreateTodo_ShouldRunSuccesfully(t *testing.T) {
	// Arrange
	cleanDb()

	requestDto := CreateRequestDto{Title: "test title", Completed: false}
	requestBody, err := json.Marshal(requestDto)
	if err != nil {
		t.Fatal(err)
	}
	
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/todos", bytes.NewBuffer(requestBody))
	w := httptest.NewRecorder()
	handler := NewHandler(dbi)

	// Act
	handler.create(w, req)

	// Assert
	resp := w.Result()
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("unexpected status code: %d", resp.StatusCode)
	}
}

func Test_Integration_GetTodos_ShouldRunSuccesfully(t *testing.T) {
	// Arrange
	cleanDb()
	
	expectedTitle := "test todo"
	_, err := dbi.Exec("insert into todos (title, completed) values ($1, $2)", expectedTitle, false)
	if err != nil {
		t.Fatal(err)
	}
	
	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/todos", nil)
	w := httptest.NewRecorder()
	handler := NewHandler(dbi)

	// Act
	handler.getAll(w, req)

	// Assert
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var todos []Model
	err = json.Unmarshal(responseBody, &todos)
	if err != nil {
		t.Fatal(err)
	}

	if len(todos) == 0 {
		t.Errorf("missing todos in response: %v", todos)
	}

	todo := todos[0]
	if todo.ID == 0 {
		t.Errorf("missing 'id' in response: %v", todo)
	}

	if todo.Title != expectedTitle {
		t.Errorf("unexpected 'title' in response: %v", todo)
	}
}

func cleanDb() {
	dbi.Exec("delete from todos")
}