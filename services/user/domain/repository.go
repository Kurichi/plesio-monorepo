package domain

import "context"

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	UpdateUser(ctx context.Context, user *User) (*User, error)
	GetUserByID(ctx context.Context, userID string) (*User, error)
	GetUsersByIDs(ctx context.Context, userIDs []string) ([]*User, error)
	GetUserByGithubID(ctx context.Context, githubID string) (*User, error)
}

type IdentityRepository interface {
	FetchUser(ctx context.Context, githubID string) (*User, error)
}
