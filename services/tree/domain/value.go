package domain

type TreeStage int

const (
	TreeStageNothing TreeStage = iota + 1
	TreeStageSeed
	TreeStageSprout
	TreeStageSapling
	TreeStageMature
	TreeStageGiantTree
)

func (ts TreeStage) Int() int {
	return int(ts)
}

func NewTreeStageFromInt(treeStage int) TreeStage {
	return TreeStage(treeStage)
}
