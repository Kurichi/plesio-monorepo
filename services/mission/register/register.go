package register

import (
	"google.golang.org/grpc"
)

func Register() *grpc.Server {
	server := grpc.NewServer()

	return server
}
