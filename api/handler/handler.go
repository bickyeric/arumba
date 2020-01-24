package handler

import "github.com/labstack/echo"

// Interface ...
type Interface interface {
	OnHandle(echo.Context) error
}
