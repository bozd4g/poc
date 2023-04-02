package db

import "database/sql"

func MigrateUp(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS todos
			(
				id SERIAL PRIMARY KEY,
				title varchar,
				completed boolean
			);
	`
	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
