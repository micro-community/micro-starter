package handler

import (
	"context"

	role "github.com/micro-community/auth/protos"
	"github.com/micro/micro/v3/service"
)

//Role implements the auth service interface
type Role struct {
	RoleID string
	Name   string
}

// NewRole returns an initUser handler
func NewRole(service *service.Service) *Role {
	return &Role{
		Name: service.Name(),
	}
}

func (e *User) GetRole(ctx context.Context, req *role.GetRoleRequest, resp *role.GetRoleResponse) error {
	panic("no implemention")
}

func (e *User) InsertRole(ctx context.Context, req *role.InsertRoleRequest, resp *role.InsertRoleResponse) error {
	panic("no implemention")
}

func (e *User) DeleteRole(ctx context.Context, req *role.DeleteRoleRequest, resp *role.DeleteRoleResponse) error {
	panic("no implemention")
}

func (e *User) UpdateRole(ctx context.Context, req *role.UpdateRoleRequest, resp *role.UpdateRoleResponse) error {
	panic("no implemention")
}
