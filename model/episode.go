package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Episode ...
type Episode struct {
	ID        primitive.ObjectID `bson:"_id"`
	ComicID   primitive.ObjectID `bson:"comic_id"`
	No        float64            `bson:"no"`
	Name      string             `bson:"name"`
	CreatedAt time.Time          `bson:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty"`
}
