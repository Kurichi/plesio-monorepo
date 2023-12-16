package infra

import (
	"context"

	"github.com/Kurichi/plesio-monorepo/services/tree/domain"
	"github.com/Kurichi/plesio-monorepo/services/tree/pkg/database"
	"github.com/uptrace/bun"
)

type treeRepositoryImpl struct {
	db *database.DB
}

func NewTreeRepository(db *database.DB) domain.TreeRepository {
	return &treeRepositoryImpl{db: db}
}

func (repo *treeRepositoryImpl) CreateTree(ctx context.Context, tree *domain.Tree) error {
	tx, err := database.GetTxByContext(ctx)
	var query *bun.InsertQuery
	if err == nil {
		query = tx.NewInsert()
	} else {
		query = repo.db.NewInsert()
	}

	modelTree := ConvertTreeModelFromEntity(tree)
	query = query.Model(modelTree)
	if _, err := query.Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (repo *treeRepositoryImpl) GetTreeByUserID(ctx context.Context, userID string) (*domain.Tree, error) {
	tx, err := database.GetTxByContext(ctx)
	var query *bun.SelectQuery
	if err == nil {
		query = tx.NewSelect()
	} else {
		query = repo.db.NewSelect()
	}

	var tree tree
	query = query.Model(&tree)
	query = query.Where("user_id = ?", userID)
	if err := query.Scan(ctx); err != nil {
		return nil, err
	}

	return tree.ConvertToEntity(), nil
}

func (repo *treeRepositoryImpl) GetTreesRanking(ctx context.Context, limit int) ([]*domain.Tree, error) {
	tx, err := database.GetTxByContext(ctx)
	var query *bun.SelectQuery
	if err == nil {
		query = tx.NewSelect()
	} else {
		query = repo.db.NewSelect()
	}

	var trees []tree
	query = query.Model(&trees).Group("user_id").OrderExpr("stage DESC", "SUM(water + fertilizer) DESC").Limit(limit)
	if err := query.Scan(ctx); err != nil {
		return nil, err
	}
	return ConvertToEntities(trees), nil
}

func (repo *treeRepositoryImpl) UpdateTree(ctx context.Context, domainTree *domain.Tree) error {
	tx, err := database.GetTxByContext(ctx)
	var query *bun.UpdateQuery
	if err == nil {
		query = tx.NewUpdate()
	} else {
		query = repo.db.NewUpdate()
	}

	tree := ConvertTreeModelFromEntity(domainTree)
	query = query.Model(tree).WherePK()
	if _, err := query.Exec(ctx); err != nil {
		return err
	}

	return nil
}
