package item

import (
	"net/http"

	itemGrpc "github.com/Kurichi/plesio-monorepo/services/item/pkg/grpc"
	"github.com/labstack/echo/v4"
)

type ItemClient struct {
	client itemGrpc.ItemServiceClient
}

func NewItemClient(client itemGrpc.ItemServiceClient) *ItemClient {
	return &ItemClient{
		client: client,
	}
}

// @Summary Get Items
// @Description Get Items
// @Tags items
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} item.GetItemsResponse
// @Failure 500 {object} string
// @Router /items [get]
func (ic *ItemClient) GetItemsHandler(c echo.Context) error {
	userID, ok := c.Get("userID").(string)
	if !ok {
		return c.JSON(http.StatusInternalServerError, "userID is not string")
	}

	ctx := c.Request().Context()
	req := &itemGrpc.GetMyInventoryRequest{
		UserId: userID,
	}
	res, err := ic.client.GetMyInventory(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	items := make([]*Item, 0, len(res.Items))
	for _, item := range res.Items {
		items = append(items, &Item{
			ID:          item.Id,
			Name:        item.Name,
			Description: item.Description,
			Quantity:    item.Quantity,
		})
	}

	return c.JSON(http.StatusOK, &GetItemsResponse{
		Items: items,
	})
}

// @Summary Use item
// @Description Use item
// @Tags items
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param item_id path string true "Item Id"
// @Success 200 {object} interface{}
// @Failure 500 {object} string
// @Router /items/{item_id} [post]
func (ic *ItemClient) UseItemHandler(c echo.Context) error {
	userID, ok := c.Get("userID").(string)
	if !ok {
		return c.JSON(http.StatusInternalServerError, "userID is not string")
	}

	itemID := c.Param("id")

	ctx := c.Request().Context()
	req := &itemGrpc.UseItemRequest{
		UserId: userID,
		ItemId: itemID,
	}
	_, err := ic.client.UseItem(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// TODO: return tree

	return c.JSON(http.StatusOK, nil)

}
