package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	subscriber "github.com/Kurichi/plesio-monorepo/services/tree/adapter/pubsub"
	"github.com/Kurichi/plesio-monorepo/services/tree/pkg/config"
	"github.com/Kurichi/plesio-monorepo/services/tree/register"
)

func main() {
	cfg := config.New()
	listener, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		panic(err)
	}

	sub, err := subscriber.NewSubscriber(context.Background())
	if err != nil {
		log.Println(err)
	}
	err = sub.Subscribe(context.Background())
	if err != nil {
		log.Println(err)
	}

	s := register.New()
	go func() {
		log.Printf("start gRPC server! port: %v", cfg.Port)
		s.Serve(listener)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}
