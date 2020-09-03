package main

import (
	"context"
	"time"

	"github.com/crazybber/user/handler"
	"github.com/crazybber/user/lib/database/global"
	user "github.com/crazybber/user/proto"

	mopentracing "github.com/micro/go-plugins/wrapper/trace/opentracing"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-micro/v3/logger"
	opentracing "github.com/opentracing/opentracing-go"
)

func main() {

	global.Init()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	service := grpc.NewService(
		micro.Name("auth.service"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
		micro.Context(ctx),
		micro.WrapHandler(mopentracing.NewHandlerWrapper(opentracing.GlobalTracer())),
	)

	service.Init()

	user.RegisterAuthDemoHandler(&handler.User{})

	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
