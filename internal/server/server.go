package server

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Start() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	registerRoutes(e)

	e.Logger.Fatal((e.Start(os.Getenv("SERVER_PORT"))))
}
