package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Kurichi/plesio-monorepo/services/bff/pkg/config"
	hellopb "github.com/Kurichi/plesio-monorepo/services/bff/pkg/grpc"
	"github.com/labstack/echo/v4"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	client hellopb.TreeServiceClient
)

func main() {
	cfg := config.New()

	// 2. gRPCサーバーとのコネクションを確立
	conn, err := grpc.Dial(
		cfg.TreeServiceAddr,

		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal("Connection failed.")
		return
	}
	defer conn.Close()

	// 3. gRPCクライアントを生成
	client = hellopb.NewTreeServiceClient(conn)

	e := echo.New()
	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
	e.GET("/tree", func(c echo.Context) error {
		req := &hellopb.GetTreeRequest{
			Id: "1",
		}
		ctx := c.Request().Context()
		res, err := client.GetTree(ctx, req)
		if err != nil {
			return c.String(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, res.GetTree())
	})

	e.Logger.Fatal(e.Start(":" + cfg.Port))
}

func Hello(ctx context.Context) {
	req := &hellopb.GetTreeRequest{
		Id: "1",
	}
	res, err := client.GetTree(ctx, req)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res.GetTree())
	}
}
