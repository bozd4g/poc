package todo

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type (
	Handler struct {
		db *sql.DB
	}

	Model struct {
		ID        int32  `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}

	CreateRequestDto struct {
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
)

func NewHandler(db *sql.DB) *Handler {
	return &Handler{db}
}

func (h *Handler) Init() {
	http.HandleFunc("/api/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.getAll(w, r)
			break
		case http.MethodPost:
			h.create(w, r)
			break
		default:
			http.Error(w, "not allowed", http.StatusMethodNotAllowed)
		}
	})
}

func (h *Handler) getAll(w http.ResponseWriter, r *http.Request) {
	rows, err := h.db.Query("select id, title, completed from todos")
	if err != nil && err != sql.ErrNoRows {
		log.Println("failed to select todos", err)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("an error occured"))
		return
	}
	defer rows.Close()

	todos := make([]Model, 0)
	for rows.Next() {
		todo := Model{}
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed)
		if err != nil {
			log.Fatal(err)
		}

		todos = append(todos, todo)
	}

	data, _ := json.Marshal(todos)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	var requestDto CreateRequestDto

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestDto)
	if err != nil {
		panic(err)
	}

	query := "insert into todos (title, completed) values ($1, $2)"
	_, err = h.db.Exec(query, requestDto.Title, requestDto.Completed)
	if err != nil {
		log.Println("failed to create todo", err)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("an error occured"))
		return
	}

	w.WriteHeader(http.StatusCreated)
}
