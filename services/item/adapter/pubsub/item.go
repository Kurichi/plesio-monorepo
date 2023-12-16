package pubsub

import (
	"context"
	"encoding/json"
	"log"

	pb "cloud.google.com/go/pubsub"
	"github.com/Kurichi/plesio-monorepo/services/item/application"
)

type (
	Event struct {
		UserID  string   `json:"userId"`
		Rewards []Reward `json:"rewards"`
	}

	Reward struct {
		ItemID string `json:"itemId"`
		Amount int    `json:"amount"`
	}
)

type ItemController struct {
	uc application.ItemUsecase
}

func NewItemController(uc application.ItemUsecase) *ItemController {
	return &ItemController{
		uc: uc,
	}
}

func (ctrl *ItemController) GetMessage(ctx context.Context, msg *pb.Message) {
	msg.Ack()
	log.Println("Received message:", string(msg.Data))

	req := &Event{}
	if err := json.Unmarshal(msg.Data, req); err != nil {
		log.Println(err)
		return
	}

	log.Println(req)

	rewards := make([]*application.RewqrdDTO, 0, len(req.Rewards))
	for _, r := range req.Rewards {
		rewards = append(rewards, &application.RewqrdDTO{
			ItemID: r.ItemID,
			Amount: r.Amount,
		})
	}
	if err := ctrl.uc.AddItem(ctx, req.UserID, rewards); err != nil {
		log.Println(err)
	}
}
