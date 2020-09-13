package handler

import (
	"context"

	"github.com/micro-community/auth/config"
	"github.com/micro-community/auth/models"
	user "github.com/micro-community/auth/protos"

	"github.com/micro/go-micro/v3/logger"
	"github.com/micro/micro/v3/service"
)

//User implements the auth service interface
type User struct {
	Name string
}

// New returns an initUser handler
func NewUser(service *service.Service) *User {
	return &User{
		Name: service.Name(),
	}
}

func NewAuth(cfg *config.Config) *User {

	return &User{}
}

func (e *User) GetUser(ctx context.Context, req *user.GetUserRequest, resp *user.UserInfo) error {
	var user models.UserModel

	user.UserId = req.UserId
	result, err := user.Get()
	if err != nil {
		logger.Error(err)
		return err
	}

	var role models.RoleModel
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

func (e *User) InsertUser(ctx context.Context, req *user.InsertUserRequest, resp *user.InsertUserResponse) error {
	var user models.UserModel

	id, err := user.Insert()
	if err != nil {
		logger.Error(err)
		return err
	}

	resp.UserId = id

	return nil
}

func (e *User) DeleteUser(ctx context.Context, req *user.DeleteUserRequest, resp *user.DeleteUserResponse) error {
	panic("implement me")
}

func (e *User) UpdateUser(ctx context.Context, req *user.UpdateUserRequest, resp *user.UserInfo) error {
	var data models.UserModel
	result, err := data.Update(req.UserId)
	if err != nil {
		logger.Error(err)
		return err
	}

	resp = result.ToView()

	return nil
}
