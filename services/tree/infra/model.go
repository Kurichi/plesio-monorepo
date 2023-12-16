package infra

import (
	"context"

	"github.com/Kurichi/plesio-monorepo/services/tree/domain"
	"github.com/Kurichi/plesio-monorepo/services/tree/pkg/database"
	"github.com/uptrace/bun"
)

type tree struct {
	bun.BaseModel `bun:"trees"`

	UserID     string `bun:",pk"`
	Stage      int
	Water      int
	Fertilizer int
	PlantAt    int64
}

func (t *tree) ConvertToEntity() *domain.Tree {
	return &domain.Tree{
		UserID:     t.UserID,
		Stage:      domain.NewTreeStageFromInt(t.Stage),
		Water:      t.Water,
		Fertilizer: t.Fertilizer,
		PlantAt:    t.PlantAt,
	}
}

func ConvertToEntities(ts []tree) []*domain.Tree {
	trees := make([]*domain.Tree, 0, len(ts))
	for _, t := range ts {
		trees = append(trees, (&t).ConvertToEntity())
	}
	return trees
}

func ConvertTreeModelFromEntity(t *domain.Tree) *tree {
	return &tree{
		UserID:     t.UserID,
		Stage:      t.Stage.Int(),
		Water:      t.Water,
		Fertilizer: t.Fertilizer,
		PlantAt:    t.PlantAt,
	}
}

func Migrate(db *database.DB) error {
	if _, err := db.NewCreateTable().Model(&tree{}).Exec(context.Background()); err != nil {
		return err
	}
	return nil
}
