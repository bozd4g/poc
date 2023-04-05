package todo

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_GetAll_WhenQueryReturnsError_ShouldReturnInternalServerError(t *testing.T) {
	// Arrange
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("select id, title, completed from todos").WillReturnError(errors.New("an error"))
	handler := &Handler{db}

	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/todos", nil)

	// Act
	handler.getAll(recorder, req)

	// Assert
	if recorder.Code != http.StatusInternalServerError {
		t.Errorf("unexpected status code: %d", recorder.Code)
	}

	if recorder.Body.String() != `an error occured` {
		t.Errorf("unexpected body: %s", recorder.Body.String())
	}
}

func Test_GetAll_ShouldReturnStatusOK(t *testing.T) {
	// Arrange
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	mock.ExpectQuery("select id, title, completed from todos").
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "completed"}).
			AddRow(1, "test todo", false))
	handler := &Handler{db}

	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/todos", nil)

	// Act
	handler.getAll(recorder, req)

	// Assert
	if recorder.Code != http.StatusOK {
		t.Errorf("unexpected status code: %d", recorder.Code)
	}

	var todos []Model
	err = json.Unmarshal(recorder.Body.Bytes(), &todos)
	if err != nil {
		t.Fatal(err)
	}

	if len(todos) != 1 {
		t.Errorf("unexpected todos length: %d", len(todos))
	}

	if todos[0].Title != "test todo" {
		t.Errorf("unexpected todo title: %s", todos[0].Title)
	}
}
