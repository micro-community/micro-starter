package subscriber

import (
	"context"

	event "github.com/micro-community/auth/protos/message"
	"github.com/micro/micro/v3/service"
	log "github.com/micro/micro/v3/service/logger"
)

type Rbac struct{
	topicList []string
}

// Pub will publish message

var ev = service.NewEvent("messages")

func (r *Rbac) Publish(ctx context.Context, msg *event.Message) error {

	ev.Publish(ctx, &event.Message{
		ID:   "1",
		Body: []byte(`this is a testing async Event`),
	})
	return nil
}

func (r *Rbac) Handle(ctx context.Context, msg *event.Message) error {
	log.Info("Received message: ", msg.Body)
	return nil
}
