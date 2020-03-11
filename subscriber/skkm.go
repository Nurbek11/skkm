package subscriber

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"

	skkm "skkm/proto/skkm"
)

type Skkm struct{}

func (e *Skkm) Handle(ctx context.Context, msg *skkm.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *skkm.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
