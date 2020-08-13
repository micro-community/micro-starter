package handler

import (
	authdemo "auth-demo/proto"
	"context"
)

func (e *AuthDemo) GetRole(ctx context.Context, args *authdemo.GetRoleArgs, resp *authdemo.GetRoleResp) error {
	panic("implement me")
}

func (e *AuthDemo) InsertRole(ctx context.Context, args *authdemo.InsertRoleArgs, resp *authdemo.InsertRoleResp) error {
	panic("implement me")
}

func (e *AuthDemo) DeleteRole(ctx context.Context, args *authdemo.DeleteRoleArgs, resp *authdemo.DeleteRoleResp) error {
	panic("implement me")
}

func (e *AuthDemo) UpdateRole(ctx context.Context, args *authdemo.UpdateRoleArgs, resp *authdemo.UpdateRoleResp) error {
	panic("implement me")
}
