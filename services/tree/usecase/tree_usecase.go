package usecase

import (
	"context"

	"github.com/Kurichi/plesio-monorepo/services/tree/domain"
)

type treeUsecase struct {
	// repo
}

func NewTreeUsecase( /*repo*/ ) TreeUsecase {
	return &treeUsecase{
		// repo
	}
}

// GetTree implements TreeUsecase.
func (uc *treeUsecase) GetTree(ctx context.Context, id string) (*TreeDTO, error) {
	// tree := uc.repo.GetTreeByID(ctx, id)
	tree, err := domain.NewTree("hoge")
	if err != nil {
		return nil, err
	}

	return NewTreeDTOFromEntity(tree), nil
}
