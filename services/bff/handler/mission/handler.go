package mission

import (
	"net/http"

	itemGrpc "github.com/Kurichi/plesio-monorepo/services/item/pkg/grpc"
	missionGrpc "github.com/Kurichi/plesio-monorepo/services/mission/pkg/grpc"
	"github.com/labstack/echo/v4"
)

type MissionClient struct {
	mClient missionGrpc.MissionServiceClient
	iClient itemGrpc.ItemServiceClient
}

func NewMissionClient(mClient missionGrpc.MissionServiceClient, iClient itemGrpc.ItemServiceClient) *MissionClient {
	return &MissionClient{
		mClient: mClient,
		iClient: iClient,
	}
}

func (mc *MissionClient) GetMissions(c echo.Context) error {
	userID, ok := c.Get("userID").(string)
	if !ok {
		return c.JSON(http.StatusInternalServerError, "userID is not string")
	}

	req := &missionGrpc.GetMissionsRequest{
		UserId: userID,
	}

	ctx := c.Request().Context()
	missions, err := mc.mClient.GetMissions(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// Get all items as a set
	set := make(map[string]bool)
	for _, mission := range missions.Missions {
		for _, item := range mission.Rewards {
			set[item.ItemId] = true
		}
	}
	itemIDs := make([]string, 0, len(set))
	for itemID := range set {
		itemIDs = append(itemIDs, itemID)
	}

	req2 := &itemGrpc.GetItemsByIDsRequest{
		ItemIds: itemIDs,
	}
	res2, err := mc.iClient.GetItemsByIDs(ctx, req2)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	items := make(map[string]*Item)
	for _, item := range res2.Items {
		items[item.Id] = &Item{
			ID:          item.Id,
			Name:        item.Name,
			Description: item.Description,
			Image:       item.Image,
		}
	}

	missionsResponse := make([]*Mission, 0, len(missions.Missions))
	for _, mission := range missions.Missions {
		itemsResponse := make([]*Item, 0, len(mission.GetRewards()))
		for _, item := range mission.GetRewards() {
			itemsResponse = append(itemsResponse, items[item.GetItemId()])
		}

		missionsResponse = append(missionsResponse, &Mission{
			ID:          mission.GetId(),
			Description: mission.GetDescription(),
			Items:       itemsResponse,
		})
	}

	return c.JSON(http.StatusOK, &GetMissionsResponse{
		Missions: missionsResponse,
	})
}

func (mc *MissionClient) ProgressMission(c echo.Context) error {
	userID, ok := c.Get("userID").(string)
	if !ok {
		return c.JSON(http.StatusInternalServerError, "userID is not string")
	}

	missionID := c.Param("id")

	req := &missionGrpc.ProgressMissionRequest{
		UserId:    userID,
		MissionId: missionID,
	}

	ctx := c.Request().Context()
	_, err := mc.mClient.ProgressMission(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, nil)
}
