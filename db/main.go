package database

import (
	"database/sql"
	"fmt"
	"millennium/db/sqlc"
)

var Queries *sqlc.Queries
var DB *sql.DB

func StartDatabase() {
	dsn := "user=postgres.udzmxmtslcuqgqifbnnn password=4cAqPzOpxdhavnIj host=aws-0-us-east-1.pooler.supabase.com port=5432 dbname=postgres"

	// Open a database connection
	DB, err := sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println("Error opening database connection:", err)
	}

	err = DB.Ping()
	if err != nil {
		fmt.Println("Error pinging the database:", err)
		return
	}
	fmt.Println("Database connection is active.")
	Queries = sqlc.New(DB)
}

func CloseDatabase() {
	DB.Close()
}