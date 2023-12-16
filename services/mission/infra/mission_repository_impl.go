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
	tx, err := database.GetTxByContext(ctx)
	var query *bun.SelectQuery
	if err == nil {
		query = tx.NewSelect()
	} else {
		query = repo.db.NewSelect()
	}

	var missions []mission
	query = query.Model(&missions)
	if filter.Term == domain.TermDaily {
		query = query.Where("term = ?", "daily")
	} else if filter.Term == domain.TermWeekly {
		query = query.Where("term = ?", "weekly")
	}
	query = query.OrderExpr("RAND()").Limit(int(count))
	if err := query.Scan(ctx); err != nil {
		return nil, err
	}
	return ConvertMissionsToEntity(missions), nil
}

// GetUserMissionByID implements domain.MissionRepository.
func (repo *missionRepositoryImpl) GetUserMissionByID(ctx context.Context, userID string, missionID string) (*domain.UserMission, error) {
	tx, err := database.GetTxByContext(ctx)
	var query *bun.SelectQuery
	if err == nil {
		query = tx.NewSelect()
	} else {
		query = repo.db.NewSelect()
	}

	var userMission userMission
	query = query.Model(&userMission).Where("user_id = ?", userID).Where("mission_id = ?", missionID)

	if err := query.Scan(ctx); err != nil {
		return nil, err
	}
	return userMission.ConvertToEntity(), nil
}

// GetUserMissions implements domain.MissionRepository.
func (repo *missionRepositoryImpl) GetUserMissions(ctx context.Context, userID string, filter domain.Filter) ([]*domain.UserMission, error) {
	tx, err := database.GetTxByContext(ctx)
	var query *bun.SelectQuery
	if err == nil {
		query = tx.NewSelect()
	} else {
		query = repo.db.NewSelect()
	}

	var userMissions []userMission
	query = query.Model(&userMissions).Relation("Mission").Where("user_id = ?", userID)
	if filter.Term == domain.TermDaily {
		query.Where("missions.term = ?", "daily")
	} else if filter.Term == domain.TermWeekly {
		query.Where("missions.term = ?", "weekly")
	}

	if _, err := query.Exec(ctx); err != nil {
		return nil, err
	}
	return ConvertUserMissionsToEntity(userMissions), nil
}

// StoreMission implements domain.MissionRepository.
func (repo *missionRepositoryImpl) StoreMission(ctx context.Context, missions []*domain.Mission) error {
	tx, err := database.GetTxByContext(ctx)
	var query *bun.InsertQuery
	if err == nil {
		query = tx.NewInsert()
	} else {
		query = repo.db.NewInsert()
	}

	modelMission := ConvertMissionsFromEntity(missions)
	query = query.Model(&modelMission)
	if _, err := query.Exec(ctx); err != nil {
		return err
	}
	return nil
}

// StoreUserMissions implements domain.MissionRepository.
func (repo *missionRepositoryImpl) StoreUserMissions(ctx context.Context, userID string, missions []*domain.UserMission) error {
	tx, err := database.GetTxByContext(ctx)
	var query *bun.InsertQuery
	if err == nil {
		query = tx.NewInsert()
	} else {
		query = repo.db.NewInsert()
	}

	userMissions := make([]*userMission, 0, len(missions))
	for _, mission := range missions {
		userMissions = append(userMissions, &userMission{
			UserID:    userID,
			MissionID: mission.Mission.ID,
			Progress:  mission.Progress,
			Deadline:  mission.Deadline,
		})
	}
	query = query.Model(&userMissions)
	if _, err := query.Exec(ctx); err != nil {
		return err
	}
	return nil
}

// UpdateUserMission implements domain.MissionRepository.
func (repo *missionRepositoryImpl) UpdateUserMission(ctx context.Context, userID string, mission *domain.UserMission) error {
	tx, err := database.GetTxByContext(ctx)
	var query *bun.UpdateQuery
	if err == nil {
		query = tx.NewUpdate()
	} else {
		query = repo.db.NewUpdate()
	}

	userMission := userMission{
		UserID:    userID,
		MissionID: mission.Mission.ID,
		Progress:  mission.Progress,
		Deadline:  mission.Deadline,
	}
	query = query.Model(&userMission).WherePK()
	if _, err := query.Exec(ctx); err != nil {
		return err
	}
	return nil
}
