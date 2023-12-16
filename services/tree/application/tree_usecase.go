package application

import (
	"context"

	"github.com/Kurichi/plesio-monorepo/services/tree/domain"
)

type treeUsecase struct {
	repo domain.TreeRepository
	tx   domain.TxRepository
}

func NewTreeUsecase(repo domain.TreeRepository, tx domain.TxRepository) TreeUsecase {
	return &treeUsecase{
		repo,
		tx,
	}
}

func (uc *treeUsecase) InitTree(ctx context.Context, userID string) (*TreeDTO, error) {
	tree, err := domain.NewTree(userID)
	if err != nil {
		return nil, err
	}
	err = uc.repo.CreateTree(ctx, tree)
	if err != nil {
		return nil, err
	}
	return NewTreeFromEntity(tree), nil
}

func (uc *treeUsecase) PlantTree(ctx context.Context, userID string) (*TreeDTO, error) {
	var tree *domain.Tree
	err := uc.tx.DoInTx(ctx, func(ctx context.Context) error {
		var err error
		tree, err = uc.repo.GetTreeByUserID(ctx, userID)
		if err != nil {
			return err
		}
		if err = tree.Plant(); err != nil {
			return err
		}
		if err = uc.repo.UpdateTree(ctx, tree); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return NewTreeFromEntity(tree), nil
}

func (uc *treeUsecase) GetMyTree(ctx context.Context, userID string) (*TreeDTO, error) {
	tree, err := uc.repo.GetTreeByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return NewTreeFromEntity(tree), nil
}

func (uc *treeUsecase) GetTreeByUserID(ctx context.Context, userID string) (*TreeDTO, error) {
	tree, err := uc.repo.GetTreeByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return NewTreeFromEntity(tree), nil
}

func (uc *treeUsecase) GetTreeRanking(ctx context.Context, limit int) ([]*TreeDTO, error) {
	ranking, err := uc.repo.GetTreesRanking(ctx, limit)
	if err != nil {
		return nil, err
	}
	return NewTreesFromEntity(ranking), nil
}
