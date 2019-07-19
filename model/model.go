package model

import (
	"errors"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

var zeroTime time.Time

// MarshalTimestamp ...
func MarshalTimestamp(t time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		if t == zeroTime {
			io.WriteString(w, strconv.FormatInt(0, 10))
		} else {
			io.WriteString(w, strconv.FormatInt(t.Unix(), 10))
		}
	})
}

// UnmarshalTimestamp ...
func UnmarshalTimestamp(v interface{}) (time.Time, error) {
	if tmpStr, ok := v.(int64); ok {
		return time.Unix(tmpStr, 0), nil
	}
	return time.Time{}, errors.New("time should be a unix timestamp")
}
