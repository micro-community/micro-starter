package handler

import (
	"context"

	"github.com/micro-community/auth/config"
	"github.com/micro-community/auth/models"
	user "github.com/micro-community/auth/protos"
	"github.com/micro/go-micro/v3/logger"
	"github.com/micro/micro/v3/service"
)

//UserHandler implements the user proto interface
type UserHandler struct {
	Name string
}

// New returns an initUser handler
func NewUser(service *service.Service) *UserHandler {
	return &UserHandler{
		Name: service.Name(),
	}
}

func NewAuth(cfg *config.Config) *UserHandler {
	return &UserHandler{}
}

func (e *UserHandler) GetUser(ctx context.Context, req *user.GetUserRequest, resp *user.UserInfo) error {
	var user models.User

	user.UserId = req.UserId
	result, err := user.Get()
	if err != nil {
		logger.Error(err)
		return err
	}

	var role models.Role
	roles, err := role.Get()
	if err != nil {
		logger.Error(err)
		return err
	}

	//transfer
	_ = result
	_ = roles
	_ = resp

	return nil
}

func (e *UserHandler) InsertUser(ctx context.Context, req *user.InsertUserRequest, resp *user.InsertUserResponse) error {
	var user models.User

	id, err := user.Insert()
	if err != nil {
		logger.Error(err)
		return err
	}

	resp.UserId = id

	return nil
}

func (e *UserHandler) DeleteUser(ctx context.Context, req *user.DeleteUserRequest, resp *user.DeleteUserResponse) error {
	panic("implement me")
}

func (e *UserHandler) UpdateUser(ctx context.Context, req *user.UpdateUserRequest, resp *user.UserInfo) error {
	var data models.User
	result, err := data.Update(req.UserId)
	if err != nil {
		logger.Error(err)
		return err
	}

	resp = result.ToView()

	return nil
}
