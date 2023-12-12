package api

import (
	"context"

	treepb "github.com/Kurichi/plesio-monorepo/services/tree/pkg/grpc"
	"github.com/Kurichi/plesio-monorepo/services/tree/usecase"
)

type treeService struct {
	treepb.UnimplementedTreeServiceServer

	uc usecase.TreeUsecase
}

func NewTreeHandler(uc usecase.TreeUsecase) treepb.TreeServiceServer {
	return &treeService{
		uc: uc,
	}
}

func (s *treeService) GetTree(ctx context.Context, req *treepb.GetTreeRequest) (*treepb.GetTreeResponse, error) {
	tree, err := s.uc.GetTree(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	res := &treepb.GetTreeResponse{
		Tree: &treepb.Tree{
			Id:   tree.ID,
			Name: tree.Name,
		},
	}
	return res, nil
}
