package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Page ...
type Page struct {
	ID        primitive.ObjectID `bson:"_id"`
	EpisodeID primitive.ObjectID `bson:"episode_id"`
	SourceID  primitive.ObjectID `bson:"source_id"`
	Link      string             `bson:"link"`
	CreatedAt time.Time          `bson:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty"`
}
