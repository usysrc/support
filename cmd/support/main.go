package main

import (
	"log"
	"net/http"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/usysrc/support/internal/api"
	"github.com/usysrc/support/internal/models"
)

func main() {
	// Connect to the database
	// Get database connection configuration from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	sslMode := os.Getenv("DB_SSLMODE")

	// Construct DSN string
	dsn := "host=" + dbHost +
		" user=" + dbUser +
		" password=" + dbPassword +
		" dbname=" + dbName +
		" port=" + dbPort +
		" sslmode=" + sslMode

	// Connect to the database
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
