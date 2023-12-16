package tree

import (
	hellopb "github.com/Kurichi/plesio-monorepo/services/tree/pkg/grpc"
)

type GetTreeParam struct {
	ID string `param:"id"`
}

type GetTreeResponse struct {
	Tree *Tree `json:"tree"`
}

type GetTreeRankingResponse struct {
	Trees []*Tree `json:"trees"`
}

type Tree struct {
	ID         string `json:"id"`
	Stage      int32  `json:"stage"`
	Water      int32  `json:"water"`
	Fertilizer int32  `json:"fertilizer"`
	PlantAt    int64  `json:"plantAt"`
}

type GetTreeRankingQuery struct {
	Limit int32 `query:"limit"`
}

func NewGetTreeResponse(tree *hellopb.Tree) *GetTreeResponse {
	return &GetTreeResponse{
		Tree: &Tree{
			ID:         tree.GetId(),
			Stage:      tree.GetStage(),
			Water:      tree.GetWater(),
			Fertilizer: tree.GetFertilizer(),
			PlantAt:    tree.GetPlantAt(),
		},
	}
}

func NewInitTreeResponse(tree *hellopb.Tree) *GetTreeResponse {
	return &GetTreeResponse{
		Tree: &Tree{
			ID:         tree.GetId(),
			Stage:      tree.GetStage(),
			Water:      tree.GetWater(),
			Fertilizer: tree.GetFertilizer(),
			PlantAt:    tree.GetPlantAt(),
		},
	}
}

func NewPlantTreeResponse(tree *hellopb.Tree) *GetTreeResponse {
	return &GetTreeResponse{
		Tree: &Tree{
			ID:         tree.GetId(),
			Stage:      tree.GetStage(),
			Water:      tree.GetWater(),
			Fertilizer: tree.GetFertilizer(),
			PlantAt:    tree.GetPlantAt(),
		},
	}
}

func NewGetTreeRankingResponse(trees []*hellopb.Tree) *GetTreeRankingResponse {
	res := &GetTreeRankingResponse{}
	for _, tree := range trees {
		res.Trees = append(res.Trees, &Tree{
			ID:         tree.GetId(),
			Stage:      tree.GetStage(),
			Water:      tree.GetWater(),
			Fertilizer: tree.GetFertilizer(),
			PlantAt:    tree.GetPlantAt(),
		})
	}
	return res
}
