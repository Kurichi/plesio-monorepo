package user

import (
	userGrpc "github.com/Kurichi/plesio-monorepo/services/user/pkg/grpc"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Avatar   string `json:"avatar"`
	GithubId string `json:"github_id"`
}

type GetUserParam struct {
	ID string `json:"id"`
}

type GetUserResponse struct {
	User *User `json:"user"`
}

type GetUsersResponse struct {
	Users []*User `json:"users"`
}

type UpdateUserRequest struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

type UpdateUserResponse struct {
	User *User `json:"user"`
}

type SignUpRequest struct {
	ID string `json:"id"`
}

type SignUpResponse struct {
	User *User `json:"user"`
}

func NewGetUserResponse(user *userGrpc.User) *GetUserResponse {
	return &GetUserResponse{
		User: &User{
			ID:       user.GetId(),
			Name:     user.GetName(),
			Avatar:   user.GetAvatar(),
			GithubId: user.GetGithubId(),
		},
	}
}

func NewCreateUserResponse(user *userGrpc.User) *GetUserResponse {
	return &GetUserResponse{
		User: &User{
			ID:       user.GetId(),
			Name:     user.GetName(),
			Avatar:   user.GetAvatar(),
			GithubId: user.GetGithubId(),
		},
	}
}

func NewUpdateUserResponse(user *userGrpc.User) *UpdateUserResponse {
	return &UpdateUserResponse{
		User: &User{
			ID:       user.GetId(),
			Name:     user.GetName(),
			Avatar:   user.GetAvatar(),
			GithubId: user.GetGithubId(),
		},
	}
}
