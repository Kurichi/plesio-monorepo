package register

import (
	userpb "github.com/Kurichi/plesio-monorepo/services/user/pkg/grpc"
	"google.golang.org/grpc"
)

func Register(userHandler userpb.UserServiceServer) *grpc.Server {
	server := grpc.NewServer()
	userpb.RegisterUserServiceServer(server, userHandler)
	return server
}
