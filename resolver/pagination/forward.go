package pagination

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type forward struct {
	cursor primitive.ObjectID
	first  int
}

func (p forward) Pipelines() mongo.Pipeline { return mongo.Pipeline{} }

func validateForward(after *string, first *int) (p forward, err error) {
	p.first = defaultLimit
	if after != nil {
		p.cursor, err = primitive.ObjectIDFromHex(*after)
		if err != nil {
			return p, ErrInvalidAfterCursor
		}
	}
	if first != nil {
		if *first < 0 {
			return p, ErrNegativeFirst
		}
		p.first = *first
	}
	return p, err
}
