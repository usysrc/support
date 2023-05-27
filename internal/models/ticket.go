package models

import (
	"time"

	"gorm.io/gorm"
)

type Ticket struct {
	ID          int64     `gorm:"primary_key"`
	Title       string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	Status      string    `gorm:"not null"`
	Priority    string    `gorm:"not null"`
	CreatedAt   time.Time `gorm:"not null;default:now()"`
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

func (t *Ticket) CollectionName() string {
	return "tickets"
}
