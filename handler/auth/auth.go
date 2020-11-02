package auth

import (
	"github.com/micro-community/micro-starter/repository"
	"github.com/micro/micro/v3/service"
)

//AuthHandler implements the user proto interface
type AuthHandler struct {
	Name        string
	User     repository.IUser     // instance of the user model
	Role     repository.IRole     // instance of the role model
	Resource repository.IResource // instance of the resource model
}

func NewAuth(service *service.Service,
	user repository.IUser,
	role repository.IRole,
	resource repository.IResource) *AuthHandler {
	return &AuthHandler{
		Name:        service.Name(),
		User:     user,
		Role:     role,
		Resource: resource,
	}
}
