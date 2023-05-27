package services

import (
	"github.com/usysrc/support/internal/models"
	"gorm.io/gorm"
)

// CreateTicket creates a new ticket
func CreateTicket(ticket *models.Ticket, db *gorm.DB) error {
	// Insert the ticket into the database
	result := db.Create(ticket)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// GetTickets returns all tickets from the database
func GetTickets(db *gorm.DB) ([]models.Ticket, error) {
	// Get all tickets from the database
	var tickets []models.Ticket
	result := db.Find(&tickets)
	if result.Error != nil {
		return nil, result.Error
	}
	return tickets, nil
}
