package db

import (
	"database/sql"
	"log"
)

func Migrations(db *sql.DB) {

	createUsersTable(db)
	createPhotosTable(db)
	createCommentsTable(db)
	createSocialMediasTable(db)
}

func createUsersTable(db *sql.DB) {
	createTable := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(255) UNIQUE NOT NULL,
			email VARCHAR(255) UNIQUE NOT NULL,
			password VARCHAR(255) NOT NULL,
			age smallint NOT NULL,
			created_at timestamptz DEFAULT now(),
			updated_at timestamptz DEFAULT now()
		)
	`
	_, err = db.Exec(createTable)

	if err != nil {
		log.Fatal("Error creating users table:", err.Error())
	}
	log.Println("success creating users table")
}

func createPhotosTable(db *sql.DB) {
	createTable := `
		CREATE TABLE IF NOT EXISTS photos (
			  id SERIAL PRIMARY KEY,
			  title VARCHAR(255) NOT NULL,
			  caption TEXT NOT NULL,
			  photo_url VARCHAR(255) NOT NULL,
			  user_id SERIAL references users(id),
			  created_at timestamptz DEFAULT now(),
			  updated_at timestamptz DEFAULT now()
		)
	`
	_, err = db.Exec(createTable)

	if err != nil {
		log.Fatal("Error creating photos table:", err.Error())
	}
	log.Println("success creating photos table")

}

func createCommentsTable(db *sql.DB) {
	createTable := `
		CREATE TABLE IF NOT EXISTS comments (
			  id SERIAL PRIMARY KEY,
			  user_id SERIAL references users(id),
			  photo_id SERIAL references photos(id),
			  message TEXT NOT NULL,
			  created_at timestamptz DEFAULT now(),
			  updated_at timestamptz DEFAULT now()
		)
	`
	_, err = db.Exec(createTable)

	if err != nil {
		log.Fatal("Error creating comments table:", err.Error())
	}
	log.Println("success creating comments table")

}

func createSocialMediasTable(db *sql.DB) {
	createTable := `
		CREATE TABLE IF NOT EXISTS social_medias (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255),
			social_media_url VARCHAR(255) ,
			user_id SERIAL references users(id)
		)
	`
	_, err = db.Exec(createTable)

	if err != nil {
		log.Fatal("Error creating social_medias table:", err.Error())
	}
	log.Println("success creating social_medias table")

}
