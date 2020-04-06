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
