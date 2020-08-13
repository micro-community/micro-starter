package subscriber

import (
	"context"
	log "github.com/micro/micro/v3/service/logger"

	authdemo "auth-demo/proto"
)

type AuthDemo struct{}

func (e *AuthDemo) Handle(ctx context.Context, msg *authdemo.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *authdemo.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
