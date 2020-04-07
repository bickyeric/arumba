package pagination

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type forward struct {
	episodeNo int
	first     int
}

func (p forward) Pipelines() (pipe mongo.Pipeline) {
	if p.episodeNo > 0 {
		pipe = append(pipe, primitive.D{
			{
				Key: "$match",
				Value: primitive.M{
					"no": primitive.M{
						"$lt": p.episodeNo,
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
