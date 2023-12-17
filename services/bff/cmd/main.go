package main

import (
	"log"
	"net/http"

	"github.com/Kurichi/plesio-monorepo/services/bff/pkg/config"
	"github.com/Kurichi/plesio-monorepo/services/bff/register"
	"github.com/labstack/echo/v4"

	_ "github.com/Kurichi/plesio-monorepo/services/bff/docs"

	echoSwagger "github.com/swaggo/echo-swagger"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	cfg := config.New()

	g := register.New()

	g.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	g.GET("/swagger/*", echoSwagger.WrapHandler)

	log.Fatal(g.Start(":" + cfg.Port))
}
