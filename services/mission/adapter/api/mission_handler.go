package api

import (
	"context"

	"github.com/Kurichi/plesio-monorepo/services/mission/application"
	missionpb "github.com/Kurichi/plesio-monorepo/services/mission/pkg/grpc"
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

func (mh *missionHandler) GetMissions(ctx context.Context, req *missionpb.GetMissionsRequest) (*missionpb.GetMissionsResponse, error) {
	userMissions, err := mh.uc.GetMissions(ctx, req.GetUserId(), req.Term)
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
