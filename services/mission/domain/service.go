package domain

import (
	"context"
	"errors"
	"time"
)

type MissionService interface {
	GenerateDailyUserMission(ctx context.Context, userID string) ([]*UserMission, error)
	GenerateWeeklyUserMission(ctx context.Context, userID string) ([]*UserMission, error)
}

type missionService struct {
	repo MissionRepository
}

func NewMissionService(repo MissionRepository) MissionService {
	return &missionService{
		repo: repo,
	}
}

// GenerateDailyUserMission implements MissionService.
func (svc *missionService) GenerateDailyUserMission(ctx context.Context, userID string) ([]*UserMission, error) {
	missions, err := svc.repo.GetRandomMissions(ctx, 5, Filter{Term: TermDaily})
	if err != nil {
		return nil, err
	}

	if len(missions) == 0 {
		return nil, errors.New("mission is not exist")
	}

	// 次の日の0時までのunix時間を取得
	year, month, date := time.Now().Add(time.Hour * 24).Date()
	deadline := time.Date(year, month, date, 0, 0, 0, 0, time.Local).Unix()

	userMissions := make([]*UserMission, 0, len(missions))
	for _, mission := range missions {
		userMissions = append(userMissions, &UserMission{
			Mission:  mission,
			Progress: 0,
			Deadline: deadline,
		})
	}

	if err := svc.repo.StoreUserMissions(ctx, userID, userMissions); err != nil {
		return nil, err
	}

	return userMissions, nil
}

// GenerateWeeklyUserMission implements MissionService.
func (svc *missionService) GenerateWeeklyUserMission(ctx context.Context, userID string) ([]*UserMission, error) {
	missions, err := svc.repo.GetRandomMissions(ctx, 5, Filter{Term: TermWeekly})
	if err != nil {
		return nil, err
	}

	// 次の日曜日の0時までのunix時間を取得
	year, month, date := time.Now().AddDate(0, 0, 7-int(time.Now().Weekday())).Date()
	deadline := time.Date(year, month, date, 0, 0, 0, 0, time.Local).Unix()

	userMissions := make([]*UserMission, 0, len(missions))
	for _, mission := range missions {
		userMissions = append(userMissions, &UserMission{
			Mission:  mission,
			Progress: 0,
			Deadline: deadline,
		})
	}

	if err := svc.repo.StoreUserMissions(ctx, userID, userMissions); err != nil {
		return nil, err
	}

	return userMissions, nil
}
