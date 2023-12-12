package register

import (
	treepb "github.com/Kurichi/plesio-monorepo/services/tree/pkg/grpc"
	"google.golang.org/grpc"
)

func Register(treeHandler treepb.TreeServiceServer) *grpc.Server {
	server := grpc.NewServer()
	treepb.RegisterTreeServiceServer(server, treeHandler)

	return server
}
