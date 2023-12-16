package infra

import (
	"context"

	"github.com/Kurichi/plesio-monorepo/services/item/pkg/database"
	"github.com/Kurichi/plesio-monorepo/services/mission/domain"
	"github.com/uptrace/bun"
)

type missionRepositoryImpl struct {
	db *database.DB
}

func NewMissionRepository(db *database.DB) domain.MissionRepository {
	return &missionRepositoryImpl{db: db}
}

// GetMissions implements domain.MissionRepository.
func (repo *missionRepositoryImpl) GetMissions(ctx context.Context, filter domain.Filter) ([]*domain.Mission, error) {
	tx, err := database.GetTxByContext(ctx)
	var query *bun.SelectQuery
	if err == nil {
		query = tx.NewSelect()
	} else {
		query = repo.db.NewSelect()
	}

	var missions []*mission
	query = query.Model(&missions)
	if filter.Term.IsValid() {
		query = query.Where("? = ?", bun.Ident("term"), filter.Term)
	}
	if err := query.Scan(ctx); err != nil {
		return nil, err
	}

	ret := make([]*domain.Mission, 0, len(missions))
	for _, m := range missions {
		ret = append(ret, m.ConvertToEntity())
	}
	return ret, nil
}

// GetRandomMissions implements domain.MissionRepository.
func (repo *missionRepositoryImpl) GetRandomMissions(ctx context.Context, count uint8, filter domain.Filter) ([]*domain.Mission, error) {
	panic("unimplemented")
}

// GetUserMissions implements domain.MissionRepository.
func (repo *missionRepositoryImpl) GetUserMissions(ctx context.Context, userID string, filter domain.Filter) ([]*domain.UserMission, error) {
	panic("unimplemented")
}

// StoreMission implements domain.MissionRepository.
func (repo *missionRepositoryImpl) StoreMission(ctx context.Context, missions []*domain.Mission) error {
	panic("unimplemented")
}

// StoreUserMissions implements domain.MissionRepository.
func (repo *missionRepositoryImpl) StoreUserMissions(ctx context.Context, userID string, missions []*domain.UserMission) error {
	panic("unimplemented")
}

// UpdateUserMission implements domain.MissionRepository.
func (repo *missionRepositoryImpl) UpdateUserMission(ctx context.Context, userID string, mission *domain.UserMission) error {
	panic("unimplemented")
}
