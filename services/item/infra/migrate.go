package infra

import (
	"context"

	"github.com/Kurichi/plesio-monorepo/services/item/pkg/database"
)

func Migrate(db *database.DB) error {
	ctx := context.Background()
	if _, err := db.NewCreateTable().Model((*item)(nil)).IfNotExists().Exec(ctx); err != nil {
		return err
	}
	if _, err := db.NewCreateTable().Model((*inventory)(nil)).IfNotExists().Exec(ctx); err != nil {
		return err
	}
	return nil
}
