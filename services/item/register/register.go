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
