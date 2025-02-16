package main

import (
	"gothstarter/handlers"
	"log"
	"log/slog"
	"os"

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

	router.GET("/*", echo.WrapHandler(public()))
	router.GET("/", handlers.HandleHome)

	listenAddr := os.Getenv("LISTEN_ADDR")
	slog.Info("HTTP server started", "listenAddr", listenAddr)
	router.Logger.Fatal(router.Start(listenAddr))
}
