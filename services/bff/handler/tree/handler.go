package tree

import (
	"net/http"

	treeGrpc "github.com/Kurichi/plesio-monorepo/services/tree/pkg/grpc"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc/metadata"
)

type TreeClient struct {
	client treeGrpc.TreeServiceClient
}

func NewTreeClient(client treeGrpc.TreeServiceClient) *TreeClient {
	return &TreeClient{
		client: client,
	}
}

// @Summary Get My Tree
// @Description GetMy Tree
// @Tags trees
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} tree.GetTreeResponse
// @Failure 500 {object} string
// @Router /trees/me [get]
func (tc *TreeClient) GetMyTreeHandler(c echo.Context) error {
	token, ok := c.Get("token").(string)
	if !ok {
		return c.String(echo.ErrInternalServerError.Code, "token not found")
	}

	userId, ok := c.Get("userId").(string)
	if !ok {
		return c.String(echo.ErrInternalServerError.Code, "userId not found")
	}

	ctx := c.Request().Context()
	md := metadata.New(map[string]string{"token": token})
	ctx = metadata.NewOutgoingContext(ctx, md)
	res, err := tc.client.GetMyTree(ctx, &treeGrpc.GetTreeRequest{
		Id: userId,
	})
	if err != nil {
		return c.String(echo.ErrInternalServerError.Code, err.Error())
	}

	resData := NewGetTreeResponse(res.GetTree())

	return c.JSON(http.StatusOK, resData)
}

// @Summary Get My Tree
// @Description GetMy Tree
// @Tags trees
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} tree.GetTreeResponse
// @Failure 500 {object} string
// @Router /trees/init [get]
func (tc *TreeClient) InitTreeHandler(c echo.Context) error {
	token, ok := c.Get("token").(string)
	if !ok {
		return c.String(echo.ErrInternalServerError.Code, "token not found")
	}

	userId, ok := c.Get("userId").(string)
	if !ok {
		return c.String(echo.ErrInternalServerError.Code, "userId not found")
	}

	ctx := c.Request().Context()
	md := metadata.New(map[string]string{"token": token})
	ctx = metadata.NewOutgoingContext(ctx, md)
	res, err := tc.client.InitTree(ctx, &treeGrpc.PlantTreeRequest{
		Id: userId,
	})
	if err != nil {
		return c.String(echo.ErrInternalServerError.Code, err.Error())
	}

	resData := NewInitTreeResponse(res.GetTree())

	return c.JSON(http.StatusOK, resData)
}

// @Summary Plant tree
// @Description Plant tree
// @Tags trees
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} tree.GetTreeResponse
// @Failure 500 {object} string
// @Router /trees/plant [post]
func (tc *TreeClient) PlantTreeHandler(c echo.Context) error {
	token, ok := c.Get("token").(string)
	if !ok {
		return c.String(echo.ErrInternalServerError.Code, "token not found")
	}

	userId, ok := c.Get("userId").(string)
	if !ok {
		return c.String(echo.ErrInternalServerError.Code, "userId not found")
	}

	ctx := c.Request().Context()
	md := metadata.New(map[string]string{"token": token})
	ctx = metadata.NewOutgoingContext(ctx, md)
	res, err := tc.client.PlantTree(ctx, &treeGrpc.PlantTreeRequest{
		Id: userId,
	})
	if err != nil {
		return c.String(echo.ErrInternalServerError.Code, err.Error())
	}

	resData := NewPlantTreeResponse(res.GetTree())

	return c.JSON(http.StatusOK, resData)
}

// @Summary Get Tree By User ID
// @Description Get Tree By User ID
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param userId path string true "User ID"
// @Success 200 {object} tree.GetTreeResponse
// @Failure 500 {object} string
// @Router /users/{userId}/tree [get]
func (tc *TreeClient) GetTreeByUserId(c echo.Context) error {
	param := &GetTreeParam{}
	if err := c.Bind(param); err != nil {
		return c.String(echo.ErrInternalServerError.Code, err.Error())
	}

	ctx := c.Request().Context()
	res, err := tc.client.GetTreeByUserId(ctx, &treeGrpc.GetTreeRequest{
		Id: param.ID,
	})
	if err != nil {
		return c.String(echo.ErrInternalServerError.Code, err.Error())
	}

	resData := NewGetTreeResponse(res.GetTree())

	return c.JSON(http.StatusOK, resData)
}

// @Summary Get Tree Ranking
// @Description Get Tree Ranking
// @Tags trees
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param limit query int false "Limit"
// @Success 200 {object} tree.GetTreeRankingResponse
// @Failure 500 {object} string
// @Router /trees/ranking [get]
func (tc *TreeClient) GetTreeRanking(c echo.Context) error {
	query := &GetTreeRankingQuery{Limit: 10}
	if err := c.Bind(query); err != nil {
		return c.String(echo.ErrInternalServerError.Code, err.Error())
	}

	ctx := c.Request().Context()
	res, err := tc.client.GetTreeRanking(ctx, &treeGrpc.GetTreeRankingRequest{
		Limit: query.Limit,
	})
	if err != nil {
		return c.String(echo.ErrInternalServerError.Code, err.Error())
	}

	resData := NewGetTreeRankingResponse(res.GetRanking())

	return c.JSON(http.StatusOK, resData)
}
