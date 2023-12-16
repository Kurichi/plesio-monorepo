package domain

import (
	"errors"
	"time"
)

type Tree struct {
	UserID     string
	Stage      TreeStage
	Water      int
	Fertilizer int
	PlantAt    int64
}

func NewTree(userID string) (*Tree, error) {
	return &Tree{
		UserID:  userID,
		Stage:   TreeStageNothing,
		PlantAt: time.Now().Unix(),
	}, nil
}

func (t *Tree) Plant() error {
	if t.Stage != TreeStageNothing {
		return errors.New("tree is already planted")
	}
	t.Stage = TreeStageSeed
	return nil
}
