package main

import (
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/Kurichi/plesio-monorepo/services/item/pkg/config"
	"github.com/Kurichi/plesio-monorepo/services/item/register"
)

func main() {
	cfg := config.New()
	listener, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		panic(err)
	}

	server := register.New()
	s := server.Server
	go func() {
		log.Printf("start gRPC server!! port: %v", cfg.Port)
		if err := s.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()
	pb := server.Subscriber
	go func() {
		if err := pb.Serve(); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}
