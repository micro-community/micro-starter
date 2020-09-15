package handler

import (
	"github.com/micro-community/auth/config"
	"github.com/micro-community/auth/service"
)

//AuthHandler implements the user proto interface
type AuthHandler struct {
	Name string
	srv  service.UserService // instance of the user service
}

func NewAuth(cfg *config.Config) *AuthHandler {
	return &AuthHandler{}
}
