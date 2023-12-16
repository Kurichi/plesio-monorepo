package register

import (
	missionpb "github.com/Kurichi/plesio-monorepo/services/mission/pkg/grpc"
	"google.golang.org/grpc"
)

func Register(missionHandler missionpb.MissionServiceServer) *grpc.Server {
	server := grpc.NewServer()
	missionpb.RegisterMissionServiceServer(server, missionHandler)

	return server
}
