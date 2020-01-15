package middleware

import (
	"github.com/bickyeric/arumba/api"
	"github.com/labstack/echo"
)

func ErrorHandler(next echo.HandlerFunc) echo.HandlerFunc {
	fn := func(c echo.Context) error {
		err := next(c)
		if err == nil {
			return nil
		}

		clientError, ok := err.(api.ClientError)
		if !ok {
			return err
		}

		status, headers := clientError.ResponseHeaders() // Get http status code and headers.
		for k, v := range headers {
			c.Response().Header().Set(k, v)
		}
		return c.JSON(status, clientError)
	}

	return fn
}
