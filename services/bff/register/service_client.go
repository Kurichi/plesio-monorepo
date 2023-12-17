package register

import (
	"os"

	itempb "github.com/Kurichi/plesio-monorepo/services/item/pkg/grpc"
	missionpb "github.com/Kurichi/plesio-monorepo/services/mission/pkg/grpc"
	treepb "github.com/Kurichi/plesio-monorepo/services/tree/pkg/grpc"
	userpb "github.com/Kurichi/plesio-monorepo/services/user/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewTreeClient() treepb.TreeServiceClient {
	conn, _ := grpc.Dial(
		os.Getenv("TREE_SERVICE_ADDR"),

		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)

	return treepb.NewTreeServiceClient(conn)
}

func NewUserClient() userpb.UserServiceClient {
	conn, _ := grpc.Dial(
		os.Getenv("USER_SERVICE_ADDR"),

		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)

	return userpb.NewUserServiceClient(conn)
}

func NewItemClient() itempb.ItemServiceClient {
	conn, _ := grpc.Dial(
		os.Getenv("ITEM_SERVICE_ADDR"),

		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)

	return itempb.NewItemServiceClient(conn)
}

func NewMissionClient() missionpb.MissionServiceClient {
	conn, _ := grpc.Dial(
		os.Getenv("MISSION_SERVICE_ADDR"),

		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)

	return missionpb.NewMissionServiceClient(conn)
}
