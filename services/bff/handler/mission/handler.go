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

// @Summary Get Missions
// @Description Get Missions
// @Tags missions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} mission.GetMissionsResponse
// @Failure 500 {object} string
// @Router /missions [get]
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

// @Summary Update Progress Mission
// @Description Update Progress Mission
// @Tags missions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param mission_id path string true "Mission ID"
// @Success 200 {object} interface{}
// @Failure 500 {object} string
// @Router /missions/{mission_id} [post]
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

// @Summary Create Mission
// @Description Create New Mission
// @Tags missions
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param mission body mission.CreateMission true "request param"
// @Success 200 {object} interface{}
// @Failure 500 {object} string
// @Router /missions [post]
func (mc *MissionClient) CreateMission(c echo.Context) error {
	var mission CreateMission
	if err := c.Bind(&mission); err != nil {
		return err
	}

	rewards := make([]*missionGrpc.Reward, 0, len(mission.Rewards))
	for _, reward := range mission.Rewards {
		rewards = append(rewards, &missionGrpc.Reward{
			ItemId: reward.ItemID,
			Amount: int32(reward.Amount),
		})
	}

	req := missionGrpc.CreateMissionRequest{
		Description: mission.Description,
		Target:      mission.Target,
		Amount:      int32(mission.Amount),
		Unit:        mission.Unit,
		Term:        mission.Term,
		Rewards:     rewards,
	}

	ctx := c.Request().Context()
	res, err := mc.mClient.CreateMission(ctx, &req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, &CreateMissionResponse{
		Mission: &DetailedMission{
			ID:            res.Mission.Id,
			CreateMission: mission,
		},
	})
}
