package handlers

import (
	"fmt"
	"gothstarter/views/home"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

type ViewHandler struct {
	SessionStore *sessions.CookieStore
}

func (h ViewHandler) HandleHome(c echo.Context) error {
	return Render(c, home.Index())
}

func (h ViewHandler) HandleProfile(c echo.Context) error {

	session, err := h.SessionStore.Get(c.Request(), "auth-session")
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	profile, ok := session.Values["profile"].(map[string]interface{})
	if !ok {
		return c.String(http.StatusInternalServerError, "Failed to get profile from session")
	}

	fmt.Println(profile)

	username := profile["nickname"].(string)
	email := profile["name"].(string)

	return Render(c, home.Profile(username, email))
}
