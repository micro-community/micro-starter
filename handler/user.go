package handler

import (
	"context"

	"github.com/micro-community/auth/models"
	user "github.com/micro-community/auth/protos"
	"github.com/micro-community/auth/service"
	"github.com/micro/go-micro/v3/logger"
	mservice "github.com/micro/micro/v3/service"
)

//UserHandler implements the user proto interface
type UserHandler struct {
	mService *mservice.Service
	Name     string
	srv      service.UserService // instance of the user service
}

// New returns an initUser handler
func NewUser(mservice *mservice.Service) *UserHandler {
	return &UserHandler{
		mService: mservice,
		Name:     "UserHandler",
	}
}

//GetUser return User By ID
func (u *UserHandler) GetUser(ctx context.Context, req *user.GetUserRequest, resp *user.UserInfo) error {
	var user models.User
	user.ID = req.UserId

	result, err := u.srv.Login()
	if err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func (u *UserHandler) InsertUser(ctx context.Context, req *user.InsertUserRequest, resp *user.InsertUserResponse) error {
	var user models.User

	id, err := user.Insert()
	if err != nil {
		logger.Error(err)
		return err
	}

	resp.UserId = id

	return nil
}

func (u *UserHandler) DeleteUser(ctx context.Context, req *user.DeleteUserRequest, resp *user.DeleteUserResponse) error {
	panic("implement me")
}

func (u *UserHandler) UpdateUser(ctx context.Context, req *user.UpdateUserRequest, resp *user.UserInfo) error {
	var data models.User

	result, err := data.Update(req.UserId)
	if err != nil {
		logger.Error(err)
		return err
	}

	resp = result.ToView()

	return nil
}
