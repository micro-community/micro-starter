package auth

import (
	mService "github.com/micro/micro/v3/service"
)

//AuthHandler implements the user proto interface
type AuthHandler struct {
	Name        string
	UserSrv     *service.UserService     // instance of the user service
	RoleSrv     *service.RoleService     // instance of the role service
	ResourceSrv *service.ResourceService // instance of the resource service
}

func NewAuth(service *mService.Service,
	user *service.UserService,
	role *service.RoleService,
	resource *service.ResourceService) *AuthHandler {
	return &AuthHandler{
		Name:        service.Name(),
		UserSrv:     user,
		RoleSrv:     role,
		ResourceSrv: resource,
	}
}
