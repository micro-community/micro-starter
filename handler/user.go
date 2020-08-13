package handler

import (
	"auth-demo/models"
	authdemo "auth-demo/proto"
	"context"

	"github.com/micro/go-micro/v3/logger"
	"github.com/prometheus/common/log"
)

func (e *AuthDemo) GetUser(ctx context.Context, args *authdemo.GetUserArgs, resp *authdemo.UserInfo) error {
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

func (e *AuthDemo) InsertUser(ctx context.Context, args *authdemo.InsertUserArgs, resp *authdemo.InsertUserResp) error {
	var user models.UserModel

	id, err := user.Insert()
	if err != nil {
		log.Error(err)
		return err
	}

	resp.UserId = id

	return nil
}

func (e *AuthDemo) DeleteUser(ctx context.Context, args *authdemo.DeleteUserArgs, resp *authdemo.DeleteUserResp) error {
	panic("implement me")
}

func (e *AuthDemo) UpdateUser(ctx context.Context, args *authdemo.UpdateUserArgs, resp *authdemo.UserInfo) error {
	var data models.UserModel
	result, err := data.Update(args.UserId)
	if err != nil {
		log.Error(err)
		return err
	}

	resp = result.ToView()

	return nil
}
