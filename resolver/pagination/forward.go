package pagination

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type forward struct {
	cursor primitive.ObjectID
	first  int
}

func (p forward) Pipelines() (pipe mongo.Pipeline) {
	if p.cursor != primitive.NilObjectID {
		pipe = append(pipe, primitive.D{
			{
				Key: "$match",
				Value: bson.M{
					"_id": bson.M{
						"$gt": p.cursor,
					},
				},
			},
		})
	}
	pipe = append(pipe, primitive.D{{
		Key:   "$limit",
		Value: p.first,
	}})
	return pipe
}

func validateForward(after *string, first *int) (p forward, err error) {
	if p.cursor, err = validateCursor(after); err != nil {
		return p, ErrInvalidAfterCursor
	}

	if p.first, err = validateLimit(first); err != nil {
		return p, ErrNegativeFirst
	}
	return p, err
}
