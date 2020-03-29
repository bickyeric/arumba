package pagination

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
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
