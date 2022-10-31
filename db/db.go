package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

func InitializeDB() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}

	dbDriver := os.Getenv("DB_DRIVER")
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	host := os.Getenv("HOST")
	database := os.Getenv("DATABASE")
	PORT := os.Getenv("PORT")

	// DBURL := fmt.Sprintf("%s:%s@%s:%s/%s?sslmode=disable", username, password, host, PORT, database)
	dbUrl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, PORT, database)
	// DBURL := fmt.Sprintf("mysql://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, PORT, database)

	db, err = sql.Open(dbDriver, dbUrl)

	if err != nil {
		log.Fatal("Error connecting to database:", err.Error())
	}

	Migrations(db)

	fmt.Println("Successfully connected to database")
}

func GetDB() *sql.DB {
	return db
}

func SetDB(dbVal *sql.DB) {
	db = dbVal
}
