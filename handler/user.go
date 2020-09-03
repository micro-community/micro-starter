package handler

import (
	"context"

	"github.com/crazybber/user/models"
	user "github.com/crazybber/user/proto"

	"github.com/micro/go-micro/v3/logger"
	"github.com/prometheus/common/log"
)

func (e *User) GetUser(ctx context.Context, args *user.GetUserArgs, resp *user.UserInfo) error {
	var user models.UserModel

	user.UserId = args.UserId
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

func (e *User) InsertUser(ctx context.Context, args *user.InsertUserArgs, resp *user.InsertUserResp) error {
	var user models.UserModel

	id, err := user.Insert()
	if err != nil {
		log.Error(err)
		return err
	}

	resp.UserId = id

	return nil
}

func (e *User) DeleteUser(ctx context.Context, args *user.DeleteUserArgs, resp *user.DeleteUserResp) error {
	panic("implement me")
}

func (e *User) UpdateUser(ctx context.Context, args *user.UpdateUserArgs, resp *user.UserInfo) error {
	var data models.UserModel
	result, err := data.Update(args.UserId)
	if err != nil {
		log.Error(err)
		return err
	}

	resp = result.ToView()

	return nil
}
