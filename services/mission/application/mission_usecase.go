package application

import (
	"context"

	"github.com/Kurichi/plesio-monorepo/services/mission/domain"
)

type missionUsecase struct {
	repo domain.MissionRepository
	svc  domain.MissionService
}

func NewMissionUsecase(repo domain.MissionRepository, svc domain.MissionService) MissionUsecase {
	return &missionUsecase{
		repo: repo,
		svc:  svc,
	}
}

// GetMyMissions implements MissionUsecase.
func (uc *missionUsecase) GetMyDailyMissions(ctx context.Context, userID string) ([]*UserMissionDTO, error) {
	dailyMissions, err := uc.repo.GetUserMissions(ctx, userID, domain.Filter{Term: domain.TermDaily})
	if err != nil {
		return nil, err
	}

	if len(dailyMissions) > 0 {
		return NewUserMissionsFromEntity(dailyMissions), nil
	}

	missions, err := uc.svc.GenerateDailyUserMission(ctx, userID)

	if err != nil {
		return nil, err
	}

	return NewUserMissionsFromEntity(missions), nil
}
