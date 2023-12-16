package infra

import (
	"context"

	"github.com/Kurichi/plesio-monorepo/services/item/domain"
	"github.com/Kurichi/plesio-monorepo/services/item/pkg/database"
	"github.com/pkg/errors"
	"github.com/uptrace/bun"
)

type inventoryRepositoryImpl struct {
	db *database.DB
}

func NewInventoryRepository(db *database.DB) domain.InventoryRepository {
	return &inventoryRepositoryImpl{db: db}
}

// GetByUserID implements domain.InventoryRepository.
func (repo *inventoryRepositoryImpl) GetByUserID(ctx context.Context, userID string) (*domain.Inventory, error) {
	tx, err := database.GetTxByContext(ctx)
	var query *bun.SelectQuery
	if err == nil {
		query = tx.NewSelect()
	} else {
		query = repo.db.NewSelect()
	}

	var inventories []*inventory
	err = query.Model(&inventories).
		Relation("Item").
		Where("? = ?", bun.Ident("user_id"), userID).
		Scan(ctx)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	items := make([]*domain.ItemWithQuantity, 0, len(inventories))
	for _, item := range inventories {
		items = append(items, item.ConvertToEntity())
	}

	return &domain.Inventory{
		UserID: userID,
		Items:  items,
	}, nil
}

// Update implements domain.InventoryRepository.
func (repo *inventoryRepositoryImpl) Update(ctx context.Context, inv *domain.Inventory) error {
	items := make([]*inventory, 0, len(inv.Items))
	for _, i := range inv.Items {
		items = append(items, &inventory{
			UserID: inv.UserID,
			Item: &item{
				ID:          i.ID.String(),
				Name:        i.Name,
				Description: i.Description,
				Target:      i.Target.String(),
				Amount:      i.Amount,
			},
			Quantity: i.Quantity,
		})
	}

	tx, err := database.GetTxByContext(ctx)
	var query *bun.UpdateQuery
	if err == nil {
		query = tx.NewUpdate().With("_data", tx.NewValues(items))
	} else {
		query = repo.db.NewUpdate().With("_data", repo.db.NewValues(items))
	}

	if _, err := query.Model((*item)(nil)).
		TableExpr("_data").
		Where("? = ?", bun.Ident("user_id"), inv.UserID).
		Exec(ctx); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// UpdateOneItem implements domain.InventoryRepository.
func (repo *inventoryRepositoryImpl) UpdateOneItem(ctx context.Context, userID string, item *domain.ItemWithQuantity) error {
	tx, err := database.GetTxByContext(ctx)
	var query *bun.UpdateQuery
	if err == nil {
		query = tx.NewUpdate()
	} else {
		query = repo.db.NewUpdate()
	}

	_, err = query.Model((*inventory)(nil)).
		Set("? = ?", bun.Ident("quantity"), item.Quantity).
		Where("? = ?", bun.Ident("user_id"), userID).
		Where("? = ?", bun.Ident("item_id"), item.ID.String()).
		Exec(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (repo *inventoryRepositoryImpl) DeleteOneItem(ctx context.Context, userID string, itemID string) error {
	tx, err := database.GetTxByContext(ctx)
	var query *bun.DeleteQuery
	if err == nil {
		query = tx.NewDelete()
	} else {
		query = repo.db.NewDelete()
	}

	_, err = query.Model((*inventory)(nil)).
		Where("? = ?", bun.Ident("user_id"), userID).
		Where("? = ?", bun.Ident("item_id"), itemID).
		Exec(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
