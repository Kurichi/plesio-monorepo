package api

import (
	"context"

	"github.com/Kurichi/plesio-monorepo/services/tree/application"
	treepb "github.com/Kurichi/plesio-monorepo/services/tree/pkg/grpc"
)

type treeService struct {
	treepb.UnimplementedTreeServiceServer

	uc application.TreeUsecase
}

func NewTreeHandler(uc application.TreeUsecase) treepb.TreeServiceServer {
	return &treeService{
		uc: uc,
	}
}

func (ts *treeService) InitTree(ctx context.Context, req *treepb.PlantTreeRequest) (*treepb.GetTreeResponse, error) {
	tree, err := ts.uc.InitTree(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &treepb.GetTreeResponse{
		Tree: convertTreeDTOToProtoTree(tree),
	}, nil
}

func (ts *treeService) PlantTree(ctx context.Context, req *treepb.PlantTreeRequest) (*treepb.GetTreeResponse, error) {
	tree, err := ts.uc.PlantTree(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &treepb.GetTreeResponse{
		Tree: convertTreeDTOToProtoTree(tree),
	}, nil
}

func (ts *treeService) GetMyTree(ctx context.Context, req *treepb.GetTreeRequest) (*treepb.GetTreeResponse, error) {
	tree, err := ts.uc.GetMyTree(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &treepb.GetTreeResponse{
		Tree: convertTreeDTOToProtoTree(tree),
	}, nil
}

func (ts *treeService) GetTreeByUserId(ctx context.Context, req *treepb.GetTreeRequest) (*treepb.GetTreeResponse, error) {
	tree, err := ts.uc.GetTreeByUserID(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &treepb.GetTreeResponse{
		Tree: convertTreeDTOToProtoTree(tree),
	}, nil
}

func (ts *treeService) GetTreeRanking(ctx context.Context, req *treepb.GetTreeRankingRequest) (*treepb.GetTreeRankingResponse, error) {
	ranking, err := ts.uc.GetTreeRanking(ctx, int(req.GetLimit()))
	if err != nil {
		return nil, err
	}
	return &treepb.GetTreeRankingResponse{
		Ranking: convertTreesDTOToProtoTrees(ranking),
	}, nil
}

// GrowthTree implements grpc.TreeServiceServer.
func (svc *treeService) GrowthTree(ctx context.Context, req *treepb.GrowthRequest) (*treepb.GrowthResponse, error) {
	dto, err := svc.uc.GrowthTree(ctx, req.GetUserId(), req.GetTarget(), int(req.GetAmount()))
	if err != nil {
		return nil, err
	}

	res := &treepb.GrowthResponse{
		Tree: &treepb.Tree{
			Id:         dto.ID,
			Stage:      int32(dto.Stage),
			Water:      int32(dto.Water),
			Fertilizer: int32(dto.Fertilizer),
			PlantAt:    dto.PlantAt,
		},
	}
	return res, nil
}

func convertTreeDTOToProtoTree(treeDto *application.TreeDTO) *treepb.Tree {
	return &treepb.Tree{
		Id:         treeDto.ID,
		Stage:      int32(treeDto.Stage),
		Water:      int32(treeDto.Water),
		Fertilizer: int32(treeDto.Fertilizer),
		PlantAt:    treeDto.PlantAt,
	}
}

func convertTreesDTOToProtoTrees(treeDtos []*application.TreeDTO) []*treepb.Tree {
	trees := make([]*treepb.Tree, 0, len(treeDtos))
	for _, treeDto := range treeDtos {
		trees = append(trees, convertTreeDTOToProtoTree(treeDto))
	}
	return trees
}
