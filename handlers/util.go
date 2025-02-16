package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"gothstarter/models"
	"net/http"

	"github.com/a-h/templ"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

func Render(c echo.Context, component templ.Component) error {
	if err := component.Render(c.Request().Context(), c.Response().Writer); err != nil {
		return err
	}

	return nil
}

func GenerateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	state := base64.StdEncoding.EncodeToString(b)
	return state, nil
}

func GetProfile(store *sessions.CookieStore, req *http.Request) *models.Profile {
	session, err := store.Get(req, "auth-session")
	if err != nil {
		return nil
	}

	profileData, ok := session.Values["profile"].(map[string]interface{})
	if !ok {
		return nil
	}

	profile, err := models.ProfileFromMap(profileData)
	if err != nil {
		return nil
	}

	return &profile
}
