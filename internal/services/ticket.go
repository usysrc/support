package services

import (
	"github.com/usysrc/support/internal/models"
	"gorm.io/gorm"
)

func CreateTicket(ticket *models.Ticket, db *gorm.DB) error {
	// Insert the ticket into the database
	result := db.Create(ticket)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
