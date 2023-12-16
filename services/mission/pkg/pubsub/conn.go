package pubsub

import (
	"context"

	ps "cloud.google.com/go/pubsub"
	"github.com/Kurichi/plesio-monorepo/services/mission/pkg/config"
)

func New(ctx context.Context, cfg *config.PusherConfig) *ps.Topic {
	client, err := ps.NewClient(ctx, cfg.ProjectID)
	if err != nil {
		return nil
	}

	return client.Topic(cfg.TopicID)
}
