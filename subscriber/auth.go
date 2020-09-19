package subscriber

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	user "github.com/micro-community/auth/protos"
)

type User struct{}

func (e *User) Handle(ctx context.Context, msg *user.Message) error {
	log.Info("Received message: ", msg.Say)
	return nil
}
