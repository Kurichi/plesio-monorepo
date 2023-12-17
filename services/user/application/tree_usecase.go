package application

import (
	"context"
	"errors"

	"github.com/Kurichi/plesio-monorepo/services/user/domain"
	"google.golang.org/grpc/metadata"
)

type userUsecase struct {
	repo domain.UserRepository
	iden domain.IdentityRepository
}

func NewUserUsecase(repo domain.UserRepository, iden domain.IdentityRepository) UserUsecase {
	return &userUsecase{
		repo,
		iden,
	}
}

// FetchUser implements UserUsecase.
func (uc *userUsecase) FetchUser(ctx context.Context, uid string) (*UserDTO, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("metadata is not found")
	}

	token := md.Get("token")
	if len(token) == 0 {
		return nil, errors.New("github_id is not found")
	}

	user, err := uc.iden.FetchUser(ctx, uid)
	if err != nil {
		return nil, err
	}

	return NewUserFromEntity(user), nil
}

// GetUserByGithubID implements UserUsecase.
func (uc *userUsecase) GetUserByGithubID(ctx context.Context, githubID string) (*UserDTO, error) {
	user, err := uc.repo.GetUserByGithubID(ctx, githubID)
	if err != nil {
		return nil, err
	}

	return NewUserFromEntity(user), nil
}

// GetUserByID implements UserUsecase.
func (uc *userUsecase) GetUserByID(ctx context.Context, userID string) (*UserDTO, error) {
	user, err := uc.repo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return NewUserFromEntity(user), nil
}

// GetUsersByIDs implements UserUsecase.
func (uc *userUsecase) GetUsersByIDs(ctx context.Context, userIDs []string) ([]*UserDTO, error) {
	if len(userIDs) == 0 {
		return nil, errors.New("userIDs is empty")
	}

	users, err := uc.repo.GetUsersByIDs(ctx, userIDs)
	if err != nil {
		return nil, err
	}

	dtoUsers := make([]*UserDTO, 0, len(users))
	for _, user := range users {
		dtoUsers = append(dtoUsers, NewUserFromEntity(user))
	}

	return dtoUsers, nil
}

// Register implements UserUsecase.
func (uc *userUsecase) Register(ctx context.Context, user *UserDTO) (*UserDTO, error) {
	if user == nil {
		return nil, errors.New("user is nil")
	}

	dbUser, err := uc.repo.GetUserByID(ctx, user.ID)
	if dbUser != nil {
		return nil, errors.New("user already exists")
	}

	domainUser, err := uc.repo.CreateUser(ctx, domain.NewUser(user.ID, user.Name, user.Avatar, user.GithubID))
	if err != nil {
		return nil, err
	}

	return NewUserFromEntity(domainUser), nil
}

// UpdateUser implements UserUsecase.
func (uc *userUsecase) UpdateUser(ctx context.Context, id string, userName string, avatar string) (*UserDTO, error) {
	prevUser, err := uc.repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	domainUser, err := uc.repo.UpdateUser(ctx, domain.NewUser(prevUser.ID, userName, avatar, prevUser.GithubID))
	if err != nil {
		return nil, err
	}

	return NewUserFromEntity(domainUser), nil
}
