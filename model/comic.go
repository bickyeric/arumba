package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Comic merepresentasikan objek komik
type Comic struct {
	ID        primitive.ObjectID `bson:"_id"`
	Name      string             `bson:"name"`
	CreatedAt time.Time          `bson:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty"`
}
