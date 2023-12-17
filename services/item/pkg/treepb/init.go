package treepb

import (
	"log"

	"github.com/Kurichi/plesio-monorepo/services/item/pkg/config"
	treegrpc "github.com/Kurichi/plesio-monorepo/services/tree/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func New(cfg *config.Config) treegrpc.TreeServiceClient {

	// gRPCサーバーとのコネクションを確立
	conn, err := grpc.Dial(
		cfg.TreeServiceAddr,

		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatal("Connection failed....")
		return nil
	}
	return treegrpc.NewTreeServiceClient(conn)
}
