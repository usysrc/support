package main

import (
	"log"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/usysrc/support/internal/api"
	"github.com/usysrc/support/internal/models"
)

func main() {
	// Connect to the database
	dsn := "host=localhost user=myuser password=mypassword dbname=mydb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Run migrations
	db.AutoMigrate(&models.Ticket{})

	// Set up the Gin router
	router := api.SetupRouter(db)

	// Start the server
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
