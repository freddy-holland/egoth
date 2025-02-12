package handlers

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Render(c echo.Context, component templ.Component) error {
	if err := component.Render(c.Request().Context(), c.Response().Writer); err != nil {
		return err
	}

	return nil
}
