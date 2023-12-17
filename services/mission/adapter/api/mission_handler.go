package api

import (
	"context"

	"github.com/Kurichi/plesio-monorepo/services/mission/application"
	missionpb "github.com/Kurichi/plesio-monorepo/services/mission/pkg/grpc"
	"google.golang.org/grpc/metadata"
)

type missionHandler struct {
	missionpb.UnimplementedMissionServiceServer

	uc application.MissionUsecase
}

func NewMissionHandler(uc application.MissionUsecase) missionpb.MissionServiceServer {
	return &missionHandler{
		uc: uc,
	}
}

// CreateMission implements grpc.MissionServiceServer.
func (mh *missionHandler) CreateMission(ctx context.Context, req *missionpb.CreateMissionRequest) (*missionpb.CreateMissionResponse, error) {
	mission, err := mh.uc.CreateMission(
		ctx, req.GetDescription(), req.GetTarget(), int(req.GetAmount()), req.GetUnit(), req.GetTerm(), convertProtoRewardsToRewardsDTO(req.Rewards),
	)
	if err != nil {
		return nil, err
	}
	return &missionpb.CreateMissionResponse{
		Mission: convertMissionDTOToProtoSimpleMission(mission),
	}, nil
}

func (mh *missionHandler) GetMissions(ctx context.Context, req *missionpb.GetMissionsRequest) (*missionpb.GetMissionsResponse, error) {
	var token string
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if len(md["token"]) > 0 {
			token = md["token"][0]
		}
	}

	userMissions, err := mh.uc.GetMissions(ctx, req.GetUserId(), req.Term, token)
	if err != nil {
		return nil, err
	}
	return &missionpb.GetMissionsResponse{
		Missions: convertUserMissionsDTOToProtoMissions(userMissions),
	}, nil
}

func (mh *missionHandler) ProgressMission(ctx context.Context, req *missionpb.ProgressMissionRequest) (*missionpb.ProgressMissionResponse, error) {
	userMission, err := mh.uc.ProgressMission(ctx, req.GetUserId(), req.GetMissionId(), int(req.GetProgress()))
	if err != nil {
		return nil, err
	}
	return &missionpb.ProgressMissionResponse{
		Mission:     convertUserMissionDTOToProtoMission(userMission),
		IsCompleted: userMission.IsCompleted,
	}, nil
}

func convertProtoRewardToRewardDTO(reward *missionpb.Reward) *application.RewardDTO {
	return &application.RewardDTO{
		ItemID: reward.ItemId,
		Amount: int(reward.Amount),
	}
}

func convertProtoRewardsToRewardsDTO(rewards []*missionpb.Reward) []*application.RewardDTO {
	rewardDTOs := make([]*application.RewardDTO, 0, len(rewards))
	for _, reward := range rewards {
		rewardDTOs = append(rewardDTOs, convertProtoRewardToRewardDTO(reward))
	}
	return rewardDTOs
}

func convertRewardDTOToProtoReward(rewardDto *application.RewardDTO) *missionpb.Reward {
	return &missionpb.Reward{
		ItemId: rewardDto.ItemID,
		Amount: int32(rewardDto.Amount),
	}
}

func convertRewardsDTOToProtoRewards(rewardDtos []*application.RewardDTO) []*missionpb.Reward {
	rewards := make([]*missionpb.Reward, 0, len(rewardDtos))
	for _, rewardDto := range rewardDtos {
		rewards = append(rewards, convertRewardDTOToProtoReward(rewardDto))
	}
	return rewards
}

func convertMissionDTOToProtoSimpleMission(missionDto *application.MissionDTO) *missionpb.SimpleMission {
	return &missionpb.SimpleMission{
		Id:          missionDto.ID,
		Description: missionDto.Description,
		Target:      missionDto.Target,
		Amount:      int32(missionDto.Amount),
		Unit:        missionDto.Unit,
		Term:        missionDto.Term,
		Rewards:     convertRewardsDTOToProtoRewards(missionDto.Rewards),
	}
}

func convertUserMissionDTOToProtoMission(userMissionDto *application.UserMissionDTO) *missionpb.Mission {
	return &missionpb.Mission{
		Id:          userMissionDto.ID,
		Description: userMissionDto.Description,
		Target:      userMissionDto.Target,
		Amount:      int32(userMissionDto.Amount),
		Unit:        userMissionDto.Unit,
		Term:        userMissionDto.Term,
		Rewards:     convertRewardsDTOToProtoRewards(userMissionDto.Rewards),
		Progress:    uint32(userMissionDto.Progress),
		Deadline:    userMissionDto.Deadline,
	}
}

func convertUserMissionsDTOToProtoMissions(userMissionDtos []*application.UserMissionDTO) []*missionpb.Mission {
	userMissions := make([]*missionpb.Mission, 0, len(userMissionDtos))
	for _, userMissionDto := range userMissionDtos {
		userMissions = append(userMissions, convertUserMissionDTOToProtoMission(userMissionDto))
	}
	return userMissions
}
