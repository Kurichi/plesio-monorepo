package application

import (
	"github.com/Kurichi/plesio-monorepo/services/user/domain"
	userpb "github.com/Kurichi/plesio-monorepo/services/user/pkg/grpc"
)

type (
	UserDTO struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Avatar   string `json:"avatar"`
		GithubID string `json:"github_id"`
	}
)

func NewUserFromEntity(e *domain.User) *UserDTO {
	return &UserDTO{
		ID:       e.ID,
		Name:     e.Name,
		Avatar:   e.Avatar,
		GithubID: e.GithubID,
	}
}

func NewUsersFromEntity(e []*domain.User) []*UserDTO {
	users := make([]*UserDTO, 0, len(e))

	for _, item := range e {
		users = append(users, NewUserFromEntity(item))
	}

	return users
}

func ConvertUserDTOToProtoUser(userDto *UserDTO) *userpb.User {
	return &userpb.User{
		Id:       userDto.ID,
		GithubId: userDto.GithubID,
		Avatar:   userDto.Avatar,
		Name:     userDto.Name,
	}
}

func ConvertUsersDTOToProtoUsers(usersDto []*UserDTO) []*userpb.User {
	users := make([]*userpb.User, 0, len(usersDto))

	for _, item := range usersDto {
		users = append(users, ConvertUserDTOToProtoUser(item))
	}

	return users
}
