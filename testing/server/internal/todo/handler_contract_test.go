//go:build provider
// +build provider

package todo

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"poc-testing/server/internal/db"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

const (
	brokerURL       = "http://localhost:9292"
	host            = "127.0.0.1"
	providerName    = "Server"
	consumerName    = "Client"
	consumerTag     = "master"
	providerVersion = "1.0.0"
	consumerVersion = "1.0.0"
)

func setup() (*sql.DB, func()) {
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

	var dbi *sql.DB
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

	return dbi, func() {
		pool.Purge(resource)
	}
}

func Test_Contract_GetTodos_ShouldRunSuccesfully(t *testing.T) {
	// Arrange
	dbi, clean := setup()
	_, err := dbi.Exec("insert into todos (title, completed) values ($1, $2)", "title", false)
	defer clean()

	handler := NewHandler(dbi)
	srv := httptest.NewServer(http.HandlerFunc(handler.getAll))
	defer srv.Close()

	pact := &dsl.Pact{
		Host:                     host,
		Consumer:                 consumerName,
		Provider:                 providerName,
		DisableToolValidityCheck: true,
		PactFileWriteMode:        "merge",
	}
	defer pact.Teardown()

	pactURL := fmt.Sprintf("%s/pacts/provider/%s/consumer/%s/version/%s.json",
		brokerURL, providerName, consumerName, consumerVersion)

	verifyRequest := types.VerifyRequest{
		BrokerURL:       brokerURL,
		ProviderBaseURL: srv.URL,
		ProviderVersion: providerVersion,
		Tags:            []string{consumerTag},
		PactURLs:        []string{pactURL},
		StateHandlers: map[string]types.StateHandler{
			"get all todos": func() error {
				return nil
			},
		},
		PublishVerificationResults: true,
		FailIfNoPactsFound:         true,
	}

	// Act & Assert
	responses, err := pact.VerifyProvider(t, verifyRequest)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%d pact tests run", len(responses))
}
