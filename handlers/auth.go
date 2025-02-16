package handlers

import (
	"fmt"
	"gothstarter/platform/authenticator"
	"net/http"
	"net/url"
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	Authenticator *authenticator.Authenticator
	SessionStore  *sessions.CookieStore
}

func (h AuthHandler) HandleLogin(c echo.Context) error {
	state, err := GenerateRandomState()
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error generating random state: %v", err.Error()))
	}

	session, _ := h.SessionStore.Get(c.Request(), "auth-session")
	session.Values["state"] = state
	err = session.Save(c.Request(), c.Response().Writer)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("Error saving session: %v", err.Error()))
	}

	return c.Redirect(http.StatusTemporaryRedirect, h.Authenticator.AuthCodeURL(state))
}

func (h AuthHandler) HandleCallback(c echo.Context) error {
	session, _ := h.SessionStore.Get(c.Request(), "auth-session")

	if c.QueryParam("state") != session.Values["state"] {
		return c.String(http.StatusBadRequest, "Invalid state parameter")
	}

	token, err := h.Authenticator.Exchange(c.Request().Context(), c.QueryParam("code"))
	if err != nil {
		return c.String(http.StatusUnauthorized, "Failed to exchange code for token")
	}

	idToken, err := h.Authenticator.VerifyIDToken(c.Request().Context(), token)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to verify ID token")
	}

	var profile map[string]interface{}
	if err := idToken.Claims(&profile); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	session.Values["access_token"] = token.AccessToken
	session.Values["profile"] = profile

	err = sessions.Save(c.Request(), c.Response().Writer)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.Redirect(http.StatusTemporaryRedirect, "/")
}

func (h AuthHandler) HandleLogout(c echo.Context) error {

	session, err := h.SessionStore.Get(c.Request(), "auth-session")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	session.Values = make(map[interface{}]interface{})
	session.Options.MaxAge = -1

	err = sessions.Save(c.Request(), c.Response().Writer)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	logoutUrl, err := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN") + "/v2/logout")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	scheme := "http"
	if c.Request().TLS != nil {
		scheme = "https"
	}

	returnTo, err := url.Parse(scheme + "://" + c.Request().Host)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	params := url.Values{}
	params.Add("returnTo", returnTo.String())
	params.Add("client_id", os.Getenv("AUTH0_CLIENT_ID"))
	logoutUrl.RawQuery = params.Encode()

	return c.Redirect(http.StatusTemporaryRedirect, logoutUrl.String())
}

func (h AuthHandler) HandleProfile(c echo.Context) error {
	session, _ := h.SessionStore.Get(c.Request(), "auth-session")
	profile := session.Values["profile"]
	accessToken := session.Values["access_token"]

	return c.JSON(http.StatusOK, map[string]interface{}{
		"profile":      profile,
		"access_token": accessToken,
	})
}
