package subscriber

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	user "github.com/crazybber/user/proto"
)

type User struct{}

func (e *User) Handle(ctx context.Context, msg *user.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *user.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
