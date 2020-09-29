package pubsub

import (
	"context"
	"errors"

	event "github.com/micro-community/auth/protos/message"
	"github.com/micro/micro/v3/service"
	log "github.com/micro/micro/v3/service/logger"
)

//rbac subscription proc for topic sub/pub
type rbac struct {
	mservice         *service.Service
	topicsPublisher  map[string]*service.Event
	topicsSubscribed []string
}

func RegisterSubscription(srv *service.Service, options *Options) {

	rbacSub := &rbac{
		mservice:        srv,
		topicsPublisher: map[string]*service.Event{},
	}

	// Add all topics to publish
	for _, topciString := range options.PubTopics {
		//topciString := "some-topic-to-publish"
		publisher := service.NewEvent(topciString)
		rbacSub.topicsPublisher[topciString] = publisher

	}

	for _, topciString := range options.SubTopics {
		//topciString := "some-topic-to-subscribe"
		srv.Subscribe(topciString, rbacSub)
	}

}

func (r *rbac) Publish(ctx context.Context, msg *event.Message) error {

	//可以发布多个主题
	if ev, found := r.topicsPublisher[msg.EventType]; found {
		ev.Publish(ctx, &event.Message{
			ID:   "1",
			Body: []byte(`this is a testing async Event`),
		})
		return nil
	}
	return errors.New("no topic exist")
}

func (r *rbac) Handle(ctx context.Context, msg *event.Message) error {
	log.Info("Received message: ", msg.Body)
	return nil
}
