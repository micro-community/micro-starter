package handler

import (
	"context"

	"github.com/crazybber/user/config"

	log "github.com/micro/micro/v3/service/logger"

	user "github.com/crazybber/user/proto"
)

type User struct {
}

func NewAuthDemo(cfg *config.Config) *User {

	return &User{}
}

// Call is a single request handler called via client.Call or the generated client code
func (e *User) Call(ctx context.Context, req *user.Request, rsp *user.Response) error {
	log.Info("Received User.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *User) Stream(ctx context.Context, req *user.StreamingRequest, stream user.AuthDemo_StreamStream) error {
	log.Infof("Received User.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&user.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *User) PingPong(ctx context.Context, stream user.AuthDemo_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&user.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
