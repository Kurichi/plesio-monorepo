package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.GET("/", func(c echo.Context) error {
		return c.String(200, "OK!")
	})
	e.GET("/health", func(c echo.Context) error {
		return c.String(200, "OK!")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
