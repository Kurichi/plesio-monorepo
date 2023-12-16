package application

import (
	"context"
)

type UserUsecase interface {
	Register(ctx context.Context, user *UserDTO) (*UserDTO, error)
	UpdateUser(ctx context.Context, id string, userName string, avatar string) (*UserDTO, error)
	GetUserByID(ctx context.Context, userID string) (*UserDTO, error)
	GetUserByGithubID(ctx context.Context, githubID string) (*UserDTO, error)
	GetUsersByIDs(ctx context.Context, userIDs []string) ([]*UserDTO, error)
	FetchUser(ctx context.Context, userId string) (*UserDTO, error)
}
