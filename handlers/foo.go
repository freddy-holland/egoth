package handlers

import (
	"github.com/labstack/echo/v4"
	"gothstarter/views/foo"
)

func HandleFoo(c echo.Context) error {
	return Render(c, foo.Index())
}
