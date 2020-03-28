package pagination

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type backward struct {
	cursor primitive.ObjectID
	last   int
}

func (p backward) Pipelines() (pipe mongo.Pipeline) {
	if p.cursor != primitive.NilObjectID {
		pipe = append(pipe, primitive.D{
			{
				Key: "$match",
				Value: bson.M{
					"_id": bson.M{
						"$lt": p.cursor,
					},
				},
			},
		})
	}
	pipe = append(pipe, primitive.D{{
		Key: "$sort",
		Value: bson.M{
			"_id": -1,
		},
	}})
	pipe = append(pipe, primitive.D{{
		Key:   "$limit",
		Value: p.last,
	}})
	pipe = append(pipe, primitive.D{{
		Key: "$sort",
		Value: bson.M{
			"_id": 1,
		},
	}})
	return pipe
}

func validateBackward(before *string, last *int) (p backward, err error) {
	p.last = defaultLimit
	if before != nil {
		p.cursor, err = primitive.ObjectIDFromHex(*before)
		if err != nil {
			return p, ErrInvalidBeforeCursor
		}
	}
	if last != nil {
		if *last < 0 {
			return p, ErrNegativeLast
		}
		p.last = *last
	}
	return p, err
}
