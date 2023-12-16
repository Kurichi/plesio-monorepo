package infra

import (
	"context"

	"github.com/Kurichi/plesio-monorepo/services/user/domain"
	"github.com/Kurichi/plesio-monorepo/services/user/pkg/database"
	"github.com/uptrace/bun"
)

type user struct {
	bun.BaseModel `bun:"users"`

	ID       string `bun:",pk"`
	Name     string
	Avatar   string
	GithubID string
}

func (u *user) ConvertToEntity() *domain.User {
	return &domain.User{
		ID:       u.ID,
		Name:     u.Name,
		Avatar:   u.Avatar,
		GithubID: u.GithubID,
	}
}

func ConvertToEntities(us []user) []*domain.User {
	users := make([]*domain.User, 0, len(us))
	for _, u := range us {
		users = append(users, (&u).ConvertToEntity())
	}
	return users
}

func ConvertUserModelFromEntity(u *domain.User) *user {
	return &user{
		ID:       u.ID,
		Name:     u.Name,
		Avatar:   u.Avatar,
		GithubID: u.GithubID,
	}
}

func Migrate(db *database.DB) error {
	if _, err := db.NewCreateTable().Model(&user{}).Exec(context.Background()); err != nil {
		return err
	}
	return nil
}
