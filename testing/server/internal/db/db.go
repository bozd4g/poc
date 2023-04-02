package db

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "postgres"
)

func Connect() (*sql.DB, error) {
	connectionString := "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"
	source := fmt.Sprintf(connectionString, host, port, user, password, dbname)
	dbi, err := sql.Open("postgres", source)
	if err != nil {
		return nil, err
	}

	if err := dbi.Ping(); err != nil {
		return nil, err
	}

	return dbi, nil
}
