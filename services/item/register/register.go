package register

import (
	itempb "github.com/Kurichi/plesio-monorepo/services/item/pkg/grpc"
	"google.golang.org/grpc"
)

func Register(
	itemCtrl itempb.ItemServiceServer,
) *grpc.Server {
	server := grpc.NewServer()

	itempb.RegisterItemServiceServer(server, itemCtrl)

	return server
}

type Server struct {
	Server     *grpc.Server
	Subscriber *pbServer
}

func NewServer(s *grpc.Server, sub *pbServer) *Server {
	return &Server{
		Server:     s,
		Subscriber: sub,
	}
}
