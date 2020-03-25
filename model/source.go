package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Source ...
type Source struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	Hostname  string             `bson:"hostname"`
	CreatedAt time.Time          `bson:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty"`
}
