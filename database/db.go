package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error is occurred  on .env file please check")
	}

	databaseUrl := os.Getenv("DATABASE_URL")

	db, errSql := sql.Open("postgres", databaseUrl)

	createStartingTables(db)

	if errSql != nil {
		fmt.Println("There is an error while connecting to the database ", err)
		panic(err)
	} else {
		Db = db
		fmt.Println("Successfully connected to database!")
	}
}

func createStartingTables(db *sql.DB) {
	// Create music table if does not exist
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS post (
    id SERIAL PRIMARY KEY,
    description TEXT,
    url TEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    enabled BOOLEAN
  )`)
	if err != nil {
		fmt.Println("Error creating table: ", err)
	}

	// Create like table if does not exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS "like" (
    id SERIAL PRIMARY KEY,
    post_id INT,
    created_at TIMESTAMP,
    FOREIGN KEY (post_id) REFERENCES post(id) ON DELETE CASCADE
  )`)
	if err != nil {
		fmt.Println("Error creating table: ", err)
	}
}
