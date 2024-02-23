package main

import (
	"fmt"
	database "millennium/db"
	"millennium/internal/app"
)

func main() {
    fmt.Println("Welcome to the POS system!")

    database.StartDatabase()
    // Start the main application
    app.Start()
}