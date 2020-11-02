package resource

import (
	"github.com/micro-community/micro-starter/repository"
	"github.com/micro/micro/v3/service"
)

//ResourceHandler implements the auth service interface
type ResourceHandler struct {
	service *service.Service
	Name    string
	repo    repository.IResource // instance of the user service
}

//Resource 资源需要定义,如：设备、物品、组织、空间、资产等
//具体的资管管理服务，需要在对应的微服务中进行实现。

// NewResource returns an initUser handler
func NewResource(service *service.Service, resource repository.IResource) *ResourceHandler {
	return &ResourceHandler{
		service: service,
		Name:    "ResourceHandler",
		repo:    resource,
	}
}
