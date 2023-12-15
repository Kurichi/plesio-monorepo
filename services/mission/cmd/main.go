package main

import (
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

	s := register.New()

	go func() {
		log.Printf("start gRPC server!! port: %v", cfg.Port)
		s.Serve(listener)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}
