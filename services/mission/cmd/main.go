package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/Kurichi/plesio-monorepo/services/mission/pkg/config"
	"github.com/Kurichi/plesio-monorepo/services/mission/register"
)

func main() {
	cfg := config.New()
	listener, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	s := register.New(ctx)

	go func() {
		log.Printf("start gRPC server!!! port: %v", cfg.Port)
		if err := s.Serve(listener); err != nil {
			panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}
