package handler

import (
	"github.com/micro-community/auth/service"
	mService "github.com/micro/micro/v3/service"
	mservice "github.com/micro/micro/v3/service"
)

//ResourceHandler implements the auth service interface
type ResourceHandler struct {
	mService *mservice.Service
	Name   string
	srv      *service.ResourceService // instance of the user service
}


//Resource 资源需要定义

// NewResource returns an initUser handler
func NewResource(service *mService.Service, resourceService *service.ResourceService) *ResourceHandler {
	return &ResourceHandler{
		Name: "ResourceHandler",
		srv: resourceService,
	}
}
