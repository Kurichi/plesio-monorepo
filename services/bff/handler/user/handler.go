package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserClient struct {
}

func NewUserClient() *UserClient {
	return &UserClient{}
}

func (h *UserClient) GetUser(c echo.Context) error {

	resData := NewGetUserResponse()

	return c.JSON(http.StatusOK, resData)
}

func (h *UserClient) CreateUser(c echo.Context) error {
	resData := NewCreateUserResponse()

	return c.JSON(http.StatusOK, resData)
}
