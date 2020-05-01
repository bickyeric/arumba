package middleware

import (
	"context"
	"errors"
	"log"

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

// Checker ...
func (a BasicAuth) Checker(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var authenticated bool
		uname, passwd, ok := c.Request().BasicAuth()
		if ok {
			authenticated = a.Username == uname && a.Password == passwd
		}
		ctx := c.Request().Context()
		ctx = context.WithValue(ctx, basicAuthCtxKey, authenticated)
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}

// IsAuthenticated ...
func (a BasicAuth) IsAuthenticated(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	log.Println("KIWKWIW")
	authenticated := ctx.Value(basicAuthCtxKey).(bool)
	if !authenticated {
		return nil, errors.New("Who are you?")
	}
	return next(ctx)
}
