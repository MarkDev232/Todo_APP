package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Connect establishes and returns a database connection.
func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/sakila")
	if err != nil {
		log.Fatal("Error opening database connection:", err)
	}

	// Ensure the connection is alive
	if err = db.Ping(); err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	fmt.Println("Connected to the database")
	return db
}
