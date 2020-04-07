package pagination

import (
	"errors"

	"github.com/bickyeric/arumba/resolver/episode"

	"go.mongodb.org/mongo-driver/mongo"
)

// ...
var (
	ErrInvalidAfterCursor = errors.New(`"after" is not valid cursor`)
	ErrNegativeFirst      = errors.New(`"first" is negative number`)

	defaultLimit = 5
)

// Interface ...
type Interface interface {
	Pipelines() mongo.Pipeline
	NextPipelines() mongo.Pipeline
}

// Validate validate pagination options from graphql
func Validate(after *string, first *int) (Interface, error) {
	var (
		p   forward
		err error
	)
	if p.episodeNo, err = validateCursor(after); err != nil {
		return p, ErrInvalidAfterCursor
	}

	if p.first, err = validateLimit(first); err != nil {
		return p, ErrNegativeFirst
	}
	return p, err
}

func validateCursor(s *string) (cursor int, err error) {
	if s != nil {
		cursor, err = episode.DecodeCursor(*s)
	}
	return cursor, err
}

func validateLimit(i *int) (limit int, err error) {
	if i == nil {
		return defaultLimit, nil
	}
	if *i < 0 {
		err = errors.New("negative value")
	}
	return *i, err
}
