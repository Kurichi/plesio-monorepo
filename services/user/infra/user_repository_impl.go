package infra

import (
	"context"

	"github.com/Kurichi/plesio-monorepo/services/user/domain"
	"github.com/Kurichi/plesio-monorepo/services/user/pkg/database"
	"github.com/uptrace/bun"
)

type userRepositoryImpl struct {
	db *database.DB
}

func NewUserRepository(db *database.DB) domain.UserRepository {
	return &userRepositoryImpl{db: db}
}

func (repo *userRepositoryImpl) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	tx, err := database.GetTxByContext(ctx)
	var query *bun.InsertQuery
	if err == nil {
		query = tx.NewInsert()
	} else {
		query = repo.db.NewInsert()
	}

	modelUser := ConvertUserModelFromEntity(user)
	query = query.Model(modelUser)
	if _, err := query.Exec(ctx); err != nil {
		return nil, err
	}

	return modelUser.ConvertToEntity(), nil
}

func (repo *userRepositoryImpl) GetUserByGithubID(ctx context.Context, githubID string) (*domain.User, error) {
	tx, err := database.GetTxByContext(ctx)
	var query *bun.SelectQuery
	if err == nil {
		query = tx.NewSelect()
	} else {
		query = repo.db.NewSelect()
	}

	var user user
	query = query.Model(&user)
	query = query.Where("github_id = ?", githubID)
	if err := query.Scan(ctx); err != nil {
		return nil, err
	}

	return user.ConvertToEntity(), nil
}

func (repo *userRepositoryImpl) GetUserByID(ctx context.Context, userID string) (*domain.User, error) {
	tx, err := database.GetTxByContext(ctx)
	var query *bun.SelectQuery
	if err == nil {
		query = tx.NewSelect()
	} else {
		query = repo.db.NewSelect()
	}

	var user user
	query = query.Model(&user)
	query = query.Where("id = ?", userID)
	if err := query.Scan(ctx); err != nil {
		return nil, err
	}

	return user.ConvertToEntity(), nil
}

func (repo *userRepositoryImpl) GetUsersByIDs(ctx context.Context, userIDs []string) ([]*domain.User, error) {
	tx, err := database.GetTxByContext(ctx)
	var query *bun.SelectQuery
	if err == nil {
		query = tx.NewSelect()
	} else {
		query = repo.db.NewSelect()
	}

	var users []user
	query = query.Model(&users)
	query = query.Where("id IN (?)", bun.In(userIDs))
	if err := query.Scan(ctx); err != nil {
		return nil, err
	}

	var entities []*domain.User
	for _, u := range users {
		entities = append(entities, u.ConvertToEntity())
	}

	return entities, nil
}

func (repo *userRepositoryImpl) UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	tx, err := database.GetTxByContext(ctx)
	var query *bun.UpdateQuery
	if err == nil {
		query = tx.NewUpdate()
	} else {
		query = repo.db.NewUpdate()
	}

	modelUser := ConvertUserModelFromEntity(user)
	query = query.Model(modelUser)
	query = query.Set("name = ?", user.Name)
	query = query.Set("avatar = ?", user.Avatar)
	query = query.Where("id = ?", user.ID)
	if _, err := query.Exec(ctx); err != nil {
		return nil, err
	}

	return modelUser.ConvertToEntity(), nil
}
