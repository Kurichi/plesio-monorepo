package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/Kurichi/plesio-monorepo/services/tree/pkg/config"
	treepb "github.com/Kurichi/plesio-monorepo/services/tree/pkg/grpc"
	"google.golang.org/grpc"
)

func main() {
	// 1. 8080番portのLisnterを作成
	cfg := config.New()
	listener, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		panic(err)
	}

	// 2. gRPCサーバーを作成
	s := grpc.NewServer()

	treepb.RegisterTreeServiceServer(s, &treeService{})

	// 3. 作成したgRPCサーバーを、8080番ポートで稼働させる
	go func() {
		log.Printf("start gRPC server port: %v", cfg.Port)
		s.Serve(listener)
	}()

	// 4.Ctrl+Cが入力されたらGraceful shutdownされるようにする
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}

type treeService struct {
	treepb.UnimplementedTreeServiceServer
}

func (s *treeService) GetTree(ctx context.Context, req *treepb.GetTreeRequest) (*treepb.GetTreeResponse, error) {
	return &treepb.GetTreeResponse{
		Tree: &treepb.Tree{
			Id:   req.GetId(),
			Name: "tree",
		},
	}, nil
}
