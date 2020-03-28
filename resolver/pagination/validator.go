package pagination

import (
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

// ...
var (
	ErrInvalidBeforeCursor = errors.New(`"before" is not valid cursor`)
	ErrInvalidAfterCursor  = errors.New(`"after" is not valid cursor`)
	ErrNegativeFirst       = errors.New(`"first" is negative number`)
	ErrNegativeLast        = errors.New(`"last" is negative number`)

	defaultLimit = 5
)

// Interface ...
type Interface interface {
	Pipelines() mongo.Pipeline
}

func Validate(before, after *string, first, last *int) (Interface, error) {
	if before != nil || last != nil {
		return validateBackward(before, last)
	} else if after != nil || first != nil {
		return validateForward(after, first)
	}
	return defaultPagination{}, nil
}

type defaultPagination struct{}

func (p defaultPagination) Pipelines() mongo.Pipeline {
	return mongo.Pipeline{
		{{
			Key:   "$limit",
			Value: defaultLimit,
		}},
	}
}
