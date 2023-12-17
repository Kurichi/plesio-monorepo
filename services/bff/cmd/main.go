package main

import (
	"log"
	"net/http"

	"github.com/Kurichi/plesio-monorepo/services/bff/pkg/config"
	"github.com/Kurichi/plesio-monorepo/services/bff/register"
	"github.com/labstack/echo/v4"

	_ "github.com/Kurichi/plesio-monorepo/services/bff/docs"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	echoSwagger "github.com/swaggo/echo-swagger"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	cfg := config.New()

	// gRPCサーバーとのコネクションを確立
	conn, err := grpc.Dial(
		cfg.TreeServiceAddr,

		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal("Connection failed...")
		return
	}
	defer conn.Close()

	g := register.New(conn)

	g.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	g.GET("/swagger/*", echoSwagger.WrapHandler)

	log.Fatal(g.Start(":" + cfg.Port))
}
