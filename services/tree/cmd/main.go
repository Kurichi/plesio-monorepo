package main

import "github.com/labstack/echo/v4"

func main() {
	e := echo.New()

	e.GET("/health", func(c echo.Context) error {
		return c.String(200, "OK!!!")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
