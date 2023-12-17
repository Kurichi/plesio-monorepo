package user

import (
	"net/http"

	treeGrpc "github.com/Kurichi/plesio-monorepo/services/tree/pkg/grpc"
	userGrpc "github.com/Kurichi/plesio-monorepo/services/user/pkg/grpc"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/metadata"
)

type UserClient struct {
	client userGrpc.UserServiceClient
	tc     treeGrpc.TreeServiceClient
}

func NewUserClient(client userGrpc.UserServiceClient, tc treeGrpc.TreeServiceClient) *UserClient {
	return &UserClient{
		client: client,
		tc:     tc,
	}
}

// @Summary Get User
// @Description Get User
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user_id path string true "user_id"
// @Success 200 {object} user.GetUserResponse
// @Failure 500 {object} string
// @Router /api/v1/users/{user_id} [get]
func (uc *UserClient) GetUser(c echo.Context) error {
	param := GetUserParam{}
	if err := c.Bind(&param); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	res, err := uc.client.GetUser(ctx, &userGrpc.GetUserRequest{
		Id: param.ID,
	})
	if err != nil {
		return c.String(echo.ErrInternalServerError.Code, err.Error())
	}

	resData := NewGetUserResponse(res.GetUser())

	return c.JSON(http.StatusOK, resData)
}

// @Summary Sign Up
// @Description Sign Up
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param body body user.SignUpRequest true "body"
// @Success 200 {object} user.SignUpResponse
// @Failure 500 {object} string
// @Router /api/v1/signup [post]
func (uc *UserClient) SignUp(c echo.Context) error {
	token, ok := c.Get("token").(string)
	if !ok {
		return c.String(echo.ErrInternalServerError.Code, "token not found")
	}

	userId, ok := c.Get("userId").(string)
	if !ok {
		return c.String(echo.ErrInternalServerError.Code, "userId not found")
	}

	body := &SignUpRequest{}
	if err := c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	md:= metadata.New(map[string]string{"token": token})
	ctx = metadata.NewOutgoingContext(ctx, md)
	user, err := uc.client.SignUp(ctx, &userGrpc.SignUpRequest{
		Id: body.ID,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	_, err = uc.tc.InitTree(ctx, &treeGrpc.PlantTreeRequest{
		Id: userId,
	})
	if err != nil {
		return c.String(echo.ErrInternalServerError.Code, err.Error())
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	resData := NewCreateUserResponse(user.GetUser())

	return c.JSON(http.StatusOK, resData)
}

// @Summary Update User
// @Description Update User
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user_id path string true "user_id"
// @Param body body user.UpdateUserRequest true "body"
// @Success 200 {object} user.UpdateUserResponse
// @Failure 500 {object} string
// @Router /api/v1/users/me [put]
func (uc *UserClient) UpdateUser(c echo.Context) error {
	userId := c.Get("userId").(string)

	body := &UpdateUserRequest{}
	if err := c.Bind(body); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := c.Request().Context()
	user, err := uc.client.UpdateUser(ctx, &userGrpc.UpdateUserRequest{
		Id:     userId,
		Name:   body.Name,
		Avatar: body.Avatar,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	resData := NewUpdateUserResponse(user.GetUser())

	return c.JSON(http.StatusOK, resData)
}
