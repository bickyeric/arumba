package pagination

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"

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

// Validate validate pagination options from graphql
func Validate(before, after *string, first, last *int) (Interface, error) {
	if before != nil || last != nil {
		return validateBackward(before, last)
	} else if first != nil {
		return validateForward(after, first)
	}
	return validateForward(after, &defaultLimit)
}

func validateCursor(s *string) (cursor primitive.ObjectID, err error) {
	if s != nil {
		cursor, err = primitive.ObjectIDFromHex(*s)
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
