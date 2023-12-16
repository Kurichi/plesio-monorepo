package api

import (
	"context"
	"log"

	"github.com/Kurichi/plesio-monorepo/services/user/application"
	userpb "github.com/Kurichi/plesio-monorepo/services/user/pkg/grpc"
)

type userService struct {
	userpb.UnimplementedUserServiceServer

	uuc application.UserUsecase
}

func NewUserHandler(uuc application.UserUsecase) userpb.UserServiceServer {
	return &userService{
		uuc: uuc,
	}
}

func (us *userService) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.GetUserResponse, error) {
	user, err := us.uuc.GetUserByID(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &userpb.GetUserResponse{
		User: application.ConvertUserDTOToProtoUser(user),
	}, nil
}

func (us *userService) GetUsers(ctx context.Context, req *userpb.GetUsersRequest) (*userpb.GetUsersResponse, error) {
	users, err := us.uuc.GetUsersByIDs(ctx, req.GetIds())
	if err != nil {
		return nil, err
	}

	log.Println("users", users)
	return &userpb.GetUsersResponse{
		Users: application.ConvertUsersDTOToProtoUsers(users),
	}, nil
}

func (us *userService) SignUp(ctx context.Context, req *userpb.SignUpRequest) (*userpb.SignUpResponse, error) {
	user, err := us.uuc.FetchUser(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	newUser, err := us.uuc.Register(ctx, user)
	if err != nil {
		return nil, err
	}

	return &userpb.SignUpResponse{
		User: application.ConvertUserDTOToProtoUser(newUser),
	}, nil

}

func (us *userService) UpdateUser(ctx context.Context, user *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	updateUser, err := us.uuc.UpdateUser(ctx, user.GetId(), user.GetName(), user.GetAvatar())
	if err != nil {
		return nil, err
	}

	return &userpb.UpdateUserResponse{
		User: application.ConvertUserDTOToProtoUser(updateUser),
	}, nil
}
