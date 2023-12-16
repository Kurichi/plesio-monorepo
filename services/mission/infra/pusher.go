package infra

import (
	"context"
	"encoding/json"

	"cloud.google.com/go/pubsub"
	"github.com/Kurichi/plesio-monorepo/services/mission/domain"
	"github.com/pkg/errors"
)

type publisher struct {
	client *pubsub.Topic
}

func NewPublisher(client *pubsub.Topic) domain.Publisher {
	return &publisher{
		client: client,
	}
}

type event struct {
	UserID  string      `json:"userId"`
	Uewards []*evReward `json:"rewards"`
}

type evReward struct {
	ItemID string `json:"itemId"`
	Amount int    `json:"amount"`
}

// Publish implements domain.Publisher.
func (pub *publisher) Publish(ctx context.Context, e *domain.CompleteMissionEvent) error {
	reward := make([]*evReward, 0, len(e.Rewards))
	for _, r := range e.Rewards {
		reward = append(reward, &evReward{
			ItemID: r.ItemID,
			Amount: r.Amount,
		})
	}

	ev := &event{
		UserID:  e.UserID,
		Uewards: reward,
	}

	data, err := json.Marshal(ev)
	if err != nil {
		return errors.WithStack(err)
	}

	res := pub.client.Publish(ctx, &pubsub.Message{
		Data: data,
	})

	<-res.Ready()
	_, err = res.Get(ctx)

	return errors.WithStack(err)
}
