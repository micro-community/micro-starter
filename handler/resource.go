package handler

import "github.com/micro/micro/v3/service"

//Resource implements the auth service interface
type ResourceHandler struct {
	RoleID string
	Name   string
}

// NewResource returns an initUser handler
func NewResource(service *service.Service) *ResourceHandler {
	return &ResourceHandler{
		Name: service.Name(),
	}
}
