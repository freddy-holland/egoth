package handlers

import (
	"gothstarter/models"
	"gothstarter/views/home"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

type ViewHandler struct {
	SessionStore *sessions.CookieStore
}

func (h ViewHandler) HandleHome(c echo.Context) error {

	var profile *models.Profile = GetProfile(h.SessionStore, c.Request())

	return Render(c, home.Index(profile))
}

func (h ViewHandler) HandleProfile(c echo.Context) error {

	var profile *models.Profile = GetProfile(h.SessionStore, c.Request())

	if profile == nil {
		return Render(c, home.Status(nil, "404", "Could not find page."))
	}

	return Render(c, home.Profile(profile))
}
