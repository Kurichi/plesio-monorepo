package application

import "context"

type MissionUsecase interface {
	CreateMission(ctx context.Context, description string, target string, amount int, unit string, term string, rewards []*RewardDTO) (*MissionDTO, error)
	GetMissions(ctx context.Context, userID string, Term *string) ([]*UserMissionDTO, error)
	ProgressMission(ctx context.Context, userID string, missionID string, progress int) (*UserMissionDTO, error)
}
