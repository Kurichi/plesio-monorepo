package domain

type TreeStage int

const (
	TreeStageNothing TreeStage = iota + 1
	TreeStageSeed
	TreeStageSprout
	TreeStageSapling
	TreeStageMature
	TreeStageGiantTree

	SproutBorderPoint    = 100
	SaplingBorderPoint   = 1000
	MatureBorderPoint    = 10000
	GiantTreeBorderPoint = 100000
)

func (ts TreeStage) Int() int {
	return int(ts)
}

func NewTreeStageFromInt(treeStage int) TreeStage {
	return TreeStage(treeStage)
}
