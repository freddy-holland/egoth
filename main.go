package main

import (
	"encoding/gob"
	"gothstarter/handlers"
	"gothstarter/platform/authenticator"
	"log"
	"log/slog"
	"os"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	router := echo.New()
	router.Use(middleware.Logger())

	auth, err := authenticator.New()
	if err != nil {
		log.Fatal(err)
	}
	store := sessions.NewCookieStore([]byte("secret"))

	gob.Register(map[string]interface{}{})

	authHandler := handlers.AuthHandler{Authenticator: auth, SessionStore: store}

	router.GET("/login", authHandler.HandleLogin)
	router.GET("/callback", authHandler.HandleCallback)
	router.GET("/logout", authHandler.HandleLogout)
	// router.GET("/profile", authHandler.HandleProfile)

	viewHandler := handlers.ViewHandler{SessionStore: store}

	router.GET("/*", echo.WrapHandler(public()))
	router.GET("/", viewHandler.HandleHome)
	router.GET("/profile", viewHandler.HandleProfile)

	listenAddr := os.Getenv("LISTEN_ADDR")
	slog.Info("HTTP server started", "listenAddr", listenAddr)
	router.Logger.Fatal(router.Start(listenAddr))
}
