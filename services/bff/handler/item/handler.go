package item

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ItemClient struct {
}

func NewItemClient() *ItemClient {
	return &ItemClient{}
}

func (ic *ItemClient) GetItemsHandler(c echo.Context) error {
	resData := NewGetItemsResponse()
	return c.JSON(http.StatusOK, resData)
}

func (ic *ItemClient) GetItemDetailHandler(c echo.Context) error {
	resData := NewGetItemResponse()
	return c.JSON(http.StatusOK, resData)
}

func (ic *ItemClient) DeleteItemHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

func (ic *ItemClient) UseItemHandler(c echo.Context) error {
	resData := NewGetItemResponse()
	return c.JSON(http.StatusOK, resData)
}
