package main

import (
	"log"
	"net/http"
	"poc-testing/server/internal/db"
	"poc-testing/server/internal/todo"

	_ "github.com/lib/pq"
)

func main() {
	dbi, err := db.Connect()
	if err != nil {
		log.Fatal("failed to connect db", err)
	}

	defer dbi.Close()
	err = db.MigrateUp(dbi)
	if err != nil {
		log.Fatal("failed to run migrations", err)
	}

	todo.NewHandler(dbi).Init()
	http.ListenAndServe(":3000", nil)
}
