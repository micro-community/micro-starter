package handler

import (
	"auth-demo/config"
	"context"

	log "github.com/micro/micro/v3/service/logger"

	authdemo "auth-demo/proto"
)

type AuthDemo struct {
}

func NewAuthDemo(cfg *config.Config) *AuthDemo {

	return &AuthDemo{}
}

// Call is a single request handler called via client.Call or the generated client code
func (e *AuthDemo) Call(ctx context.Context, req *authdemo.Request, rsp *authdemo.Response) error {
	log.Info("Received AuthDemo.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *AuthDemo) Stream(ctx context.Context, req *authdemo.StreamingRequest, stream authdemo.AuthDemo_StreamStream) error {
	log.Infof("Received AuthDemo.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&authdemo.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *AuthDemo) PingPong(ctx context.Context, stream authdemo.AuthDemo_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&authdemo.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
