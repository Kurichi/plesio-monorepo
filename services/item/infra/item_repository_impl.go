package infra

import (
	"context"

	"github.com/Kurichi/plesio-monorepo/services/item/domain"
	"github.com/Kurichi/plesio-monorepo/services/item/pkg/database"
	"github.com/pkg/errors"
	"github.com/uptrace/bun"
)

type itemRepository struct {
	db *database.DB
}

func NewItemRepository(db *database.DB) domain.ItemRepository {
	return &itemRepository{db: db}
}

// StoreItem implements domain.ItemRepository.
func (repo *itemRepository) StoreItem(ctx context.Context, i *domain.Item) error {
	tx, err := database.GetTxByContext(ctx)
	var query *bun.InsertQuery
	if err == nil {
		query = tx.NewInsert()
	} else {
		query = repo.db.NewInsert()
	}

	modelItem := item{
		ID:          i.ID.String(),
		Name:        i.Name,
		Description: i.Description,
		Target:      i.Target.String(),
		Amount:      i.Amount,
	}
	query = query.Model(&modelItem)
	if _, err := query.Exec(ctx); err != nil {
		return err
	}
	return nil
}

// GetByID implements domain.ItemRepository.
func (repo *itemRepository) GetByID(ctx context.Context, id domain.ItemID) (*domain.Item, error) {
	tx, err := database.GetTxByContext(ctx)
	var query *bun.SelectQuery
	if err == nil {
		query = tx.NewSelect()
	} else {
		query = repo.db.NewSelect()
	}

	var item item
	err = query.
		Model(&item).
		Where("? = ?", bun.Ident("id"), id).
		Scan(ctx)
	if err != nil {
		switch errors.Cause(err) {
		// TODO: add custom error for not found
		default:
			return nil, errors.WithStack(err)
		}
	}

	return item.ConvertToEntity(), nil
}

// GetByIDs implements domain.ItemRepository.
func (repo *itemRepository) GetByIDs(ctx context.Context, id []domain.ItemID) ([]*domain.Item, error) {
	tx, err := database.GetTxByContext(ctx)
	var query *bun.SelectQuery
	if err == nil {
		query = tx.NewSelect()
	} else {
		query = repo.db.NewSelect()
	}

	var items []item
	err = query.
		Model(&items).
		Where("? IN (?)", bun.Ident("id"), id).
		Scan(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	entities := make([]*domain.Item, 0, len(items))
	for _, i := range items {
		entities = append(entities, i.ConvertToEntity())
	}

	return entities, nil
}
