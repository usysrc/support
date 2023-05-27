package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Ticket struct {
	ID          bson.ObjectId `bson:"_id,omitempty"`
	Name        string        `bson:"name"`
	Description string        `bson:"description"`
	Status      string        `bson:"status"`
	Progress    int           `bson:"progress"`
	CreatedAt   time.Time     `bson:"created_at"`
	UpdatedAt   time.Time     `bson:"updated_at"`
}

func (t *Ticket) CollectionName() string {
	return "tickets"
}
