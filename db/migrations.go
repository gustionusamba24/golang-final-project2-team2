package db

import (
	"database/sql"
	"log"
)

func Migrations(db *sql.DB){

	createUserTable(db)
}

func createUserTable(db *sql.DB){
	createTable := `
		CREATE TABLE IF NOT EXISTS user (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			price NUMERIC NOT NULL,
			stock smallint NOT NULL,
			created_at timestamptz DEFAULT now()
		)
	`

	_, err = db.Exec(createTable)

	if err != nil {
		log.Fatal("Error creating user table:", err.Error())
	}
}