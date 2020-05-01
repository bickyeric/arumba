package middleware

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/labstack/echo"
)

var basicAuthCtxKey = &contextKey{name: "basic-authenticated"}

type contextKey struct {
	name string
}

// BasicAuth ...
type BasicAuth struct {
	Username, Password string
}

// Assignor ...
func (a BasicAuth) Assignor(uname string, passwd string, c echo.Context) (bool, error) {
	ctx := c.Request().Context()
	ctx = context.WithValue(ctx, basicAuthCtxKey, a.Username == uname && a.Password == passwd)
	c.SetRequest(c.Request().WithContext(ctx))

	return true, nil
}

// IsAuthenticated ...
func (a BasicAuth) IsAuthenticated(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	authenticated := ctx.Value(basicAuthCtxKey).(bool)
	if !authenticated {
		return nil, errors.New("Who are you?")
	}
	return next(ctx)
}
