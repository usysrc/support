package models

import (
	"gorm.io/gorm"
)

type Ticket struct {
	ID          string         `json:"id" bson:"_id,omitempty" gorm:"primaryKey"`
	Title       string         `json:"title" bson:"title"`
	Description string         `json:"description" bson:"description"`
	CreatedAt   gorm.DeletedAt `json:"createdAt" bson:"createdAt" gorm:"index"`
	UpdatedAt   gorm.DeletedAt `json:"updatedAt" bson:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt" bson:"deletedAt" gorm:"index"`
}

func (t *Ticket) CollectionName() string {
	return "tickets"
}
