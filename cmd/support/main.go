package main

import (
	"log"
	"net/http"

	"github.com/usysrc/support/internal/api"
	"github.com/usysrc/support/internal/database"
)

func main() {
	// Set up the Postgres database
	db := database.SetupDatabase()

	// Set up the Gin router
	router := api.SetupRouter(db)

	// Start the server
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
