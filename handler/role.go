package handler

import (
	"context"

	user "github.com/crazybber/user/proto"
)

func (e *User) GetRole(ctx context.Context, args *user.GetRoleArgs, resp *user.GetRoleResp) error {
	panic("implement me")
}

func (e *User) InsertRole(ctx context.Context, args *user.InsertRoleArgs, resp *user.InsertRoleResp) error {
	panic("implement me")
}

func (e *User) DeleteRole(ctx context.Context, args *user.DeleteRoleArgs, resp *user.DeleteRoleResp) error {
	panic("implement me")
}

func (e *User) UpdateRole(ctx context.Context, args *user.UpdateRoleArgs, resp *user.UpdateRoleResp) error {
	panic("implement me")
}
