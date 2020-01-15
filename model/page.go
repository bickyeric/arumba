package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Page ...
type Page struct {
	ID        primitive.ObjectID `json:"ID" bson:"_id"`
	EpisodeID primitive.ObjectID `json:"episodeID" bson:"episode_id"`
	SourceID  primitive.ObjectID `json:"sourceID" bson:"source_id"`
	Link      string             `json:"link" bson:"link"`
	CreatedAt time.Time          `json:"createdAt" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updated_at,omitempty"`
}
