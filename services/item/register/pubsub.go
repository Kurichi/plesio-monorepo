package register

import (
	"context"
	"log"

	pb "cloud.google.com/go/pubsub"
	"github.com/Kurichi/plesio-monorepo/services/item/adapter/pubsub"
	"github.com/Kurichi/plesio-monorepo/services/item/pkg/config"
	"github.com/pkg/errors"
)

type pbServer struct {
	client *pb.Client
	sub    *pb.Subscription
	ctrl   *pubsub.ItemController
}

func Subscriber(
	ctrl *pubsub.ItemController,
) *pbServer {
	cfg := config.NewSubscriberConfig()
	client, err := pb.NewClient(context.Background(), cfg.ProjectID)
	if err != nil {
		log.Println(err)
		return nil
	}

	s := &pbServer{
		client: client,
		sub:    client.Subscription(cfg.TopicID),
		ctrl:   ctrl,
	}

	return s
}

func (s *pbServer) Serve() error {
	for {
		if err := s.sub.Receive(context.Background(), s.ctrl.GetMessage); err != nil {
			return errors.WithStack(err)
		}
	}
}
