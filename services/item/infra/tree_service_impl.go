package infra

import (
	"context"

	"github.com/Kurichi/plesio-monorepo/services/item/domain"
	treepb "github.com/Kurichi/plesio-monorepo/services/tree/pkg/grpc"
	"github.com/pkg/errors"
)

type treeService struct {
	client treepb.TreeServiceClient
}

func NewTreeService(client treepb.TreeServiceClient) domain.TreeService {
	return &treeService{
		client: client,
	}
}

// GrowthEvent implements domain.TreeService.
func (svc *treeService) GrowthEvent(ctx context.Context, event *domain.ItemUsedEvent) (*domain.Tree, error) {
	res, err := svc.client.GrowthTree(ctx, &treepb.GrowthRequest{
		UserId: event.UserID,
		Target: event.Target.String(),
		Amount: int32(event.Amount),
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &domain.Tree{
		Stage:      int(res.Tree.Stage),
		Water:      int(res.Tree.Water),
		Fertilizer: int(res.Tree.Fertilizer),
	}, nil
}
