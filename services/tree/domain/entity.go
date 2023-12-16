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

func (t *Tree) Growth() (bool, error) {
	isStageUp := false
	switch t.Stage {
	case TreeStageSeed:
		if t.Water+t.Fertilizer >= SproutBorderPoint {
			t.Stage = TreeStageSprout
			isStageUp = true
		}
	case TreeStageSprout:
		if t.Water+t.Fertilizer >= SaplingBorderPoint {
			t.Stage = TreeStageSapling
			isStageUp = true
		}
	case TreeStageSapling:
		if t.Water+t.Fertilizer >= MatureBorderPoint {
			t.Stage = TreeStageMature
			isStageUp = true
		}
	case TreeStageMature:
		if t.Water+t.Fertilizer >= GiantTreeBorderPoint {
			t.Stage = TreeStageGiantTree
			isStageUp = true
		}
	}
	return isStageUp, nil
}
