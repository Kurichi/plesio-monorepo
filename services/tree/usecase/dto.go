package usecase

import "github.com/Kurichi/plesio-monorepo/services/tree/domain"

type TreeDTO struct {
	ID   string
	Name string
}

func NewTreeDTOFromEntity(tree *domain.Tree) *TreeDTO {
	return &TreeDTO{
		ID:   tree.ID.String(),
		Name: tree.Name,
	}
}
