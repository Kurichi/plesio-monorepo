package infra

import (
	"context"
	"log"

	"github.com/Kurichi/plesio-monorepo/services/mission/domain"
	"github.com/Kurichi/plesio-monorepo/services/mission/pkg/database"
	"github.com/uptrace/bun"
)

type mission struct {
	bun.BaseModel `bun:"missions"`

	ID          string `bun:",pk"`
	Description string
	Term        string
	Target      string
	Amount      int
	Unit        string
	Rewards     []*reward `bun:"rel:has-many,join:id=mission_id"`
}

type reward struct {
	bun.BaseModel `bun:"rewards"`

	MissionID string `bun:",pk"`
	ItemID    string `bun:",pk"`
	Amount    int
}

type userMission struct {
	bun.BaseModel `bun:"user_missions"`

	UserID    string   `bun:",pk"`
	MissionID string   `bun:",pk"`
	Mission   *mission `bun:"rel:belongs-to,join:mission_id=id"`
	Progress  int
	Deadline  int64
}

func (m *mission) ConvertToEntity() *domain.Mission {
	return &domain.Mission{
		ID:          m.ID,
		Description: m.Description,
		Term:        m.Term,
		Target:      m.Target,
		Amount:      m.Amount,
		Unit:        m.Unit,
		Rewards:     m.convertRewardsToEntity(),
	}
}

func ConvertMissionsToEntity(missions []mission) []*domain.Mission {
	domainMissions := make([]*domain.Mission, 0, len(missions))
	for _, mission := range missions {
		domainMissions = append(domainMissions, mission.ConvertToEntity())
	}
	return domainMissions
}

func (m *mission) convertRewardsToEntity() []*domain.Reward {
	rewards := make([]*domain.Reward, len(m.Rewards))
	for i, r := range m.Rewards {
		rewards[i] = r.ConvertToEntity()
	}
	return rewards
}

func (r *reward) ConvertToEntity() *domain.Reward {
	return &domain.Reward{
		ItemID: r.ItemID,
		Amount: r.Amount,
	}
}

func (um *userMission) ConvertToEntity() *domain.UserMission {
	log.Printf("%+v", um)
	return &domain.UserMission{
		UserID:   um.UserID,
		Mission:  um.Mission.ConvertToEntity(),
		Progress: um.Progress,
		Deadline: um.Deadline,
	}
}

func ConvertUserMissionsToEntity(userMissions []userMission) []*domain.UserMission {
	domainUserMissions := make([]*domain.UserMission, 0, len(userMissions))
	for _, userMission := range userMissions {
		domainUserMissions = append(domainUserMissions, userMission.ConvertToEntity())
	}
	return domainUserMissions
}

func ConvertRewardFromEntity(r *domain.Reward, missionID string) *reward {
	return &reward{
		MissionID: missionID,
		ItemID:    r.ItemID,
		Amount:    r.Amount,
	}
}

func ConvertRewardsFromEntity(r []*domain.Reward, missionID string) []*reward {
	rewards := make([]*reward, 0, len(r))
	for _, reward := range r {
		rewards = append(rewards, ConvertRewardFromEntity(reward, missionID))
	}
	return rewards
}

func ConvertMissionFromEntity(m *domain.Mission) *mission {
	return &mission{
		ID:          m.ID,
		Description: m.Description,
		Term:        m.Term,
		Target:      m.Target,
		Amount:      m.Amount,
		Unit:        m.Unit,
		Rewards:     ConvertRewardsFromEntity(m.Rewards, m.ID),
	}
}

func ConvertMissionsFromEntity(m []*domain.Mission) []*mission {
	missions := make([]*mission, 0, len(m))
	for _, mission := range m {
		missions = append(missions, ConvertMissionFromEntity(mission))
	}
	return missions
}

func Migrate(db *database.DB) error {
	if _, err := db.NewCreateTable().Model(&mission{}).Exec(context.Background()); err != nil {
		return err
	}
	if _, err := db.NewCreateIndex().Model(&mission{}).Index("term_idx").Column("term").Exec(context.Background()); err != nil {
		return err
	}
	if _, err := db.NewCreateTable().Model(&reward{}).Exec(context.Background()); err != nil {
		return err
	}
	if _, err := db.NewCreateTable().Model(&userMission{}).Exec(context.Background()); err != nil {
		return err
	}
	return nil
}
