package application

import "github.com/Kurichi/plesio-monorepo/services/tree/domain"

type (
	TreeDTO struct {
		ID         string
		UserID     string
		Stage      int
		Water      int
		Fertilizer int
		PlantAt    int64
	}
)

func NewTreeFromEntity(e *domain.Tree) *TreeDTO {
	return &TreeDTO{
		UserID:     e.UserID,
		Stage:      e.Stage.Int(),
		Water:      e.Water,
		Fertilizer: e.Fertilizer,
		PlantAt:    e.PlantAt,
	}
}

func NewTreesFromEntity(e []*domain.Tree) []*TreeDTO {
	trees := make([]*TreeDTO, 0, len(e))

	for _, item := range e {
		trees = append(trees, NewTreeFromEntity(item))
	}

	return trees
}
