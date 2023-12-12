package subscriber

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
	"github.com/Kurichi/plesio-monorepo/services/tree/pkg/config"
	"github.com/pkg/errors"
)

type Subscriber struct {
	client *pubsub.Client
	topic  *pubsub.Topic
}

func NewSubscriber(ctx context.Context) (*Subscriber, error) {
	cfg := config.NewPubSubConfig()
	client, err := pubsub.NewClient(ctx, cfg.ProjectID)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &Subscriber{
		client: client,
		topic:  client.Topic(cfg.TopicName),
	}, nil
}

func (s *Subscriber) Subscribe(ctx context.Context) error {
	sub, err := s.client.CreateSubscription(ctx, s.topic.ID(), pubsub.SubscriptionConfig{
		Topic: s.topic,
	})
	if err != nil {
		return errors.WithStack(err)
	}

	fmt.Println("start subscribe")
	err = sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		log.Println(string(m.Data))
	})
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
