package handlers

import "github.com/labstack/echo/v4"

type HTTPHandler func(c echo.Context)
