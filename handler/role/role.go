package role

import (
	"context"

	role "github.com/micro-community/micro-starter/protos"

	mservice "github.com/micro/micro/v3/service"
)

//RoleHandler implements the role proto interface
type RoleHandler struct {
	RoleID  string
	Name    string
	service *service.RoleService
}

// NewRole returns an initUser handler
func NewRole(service *mservice.Service, roleService *service.RoleService) *RoleHandler {
	return &RoleHandler{
		Name:    service.Name(),
		service: roleService,
	}
}

func (r *RoleHandler) GetRole(ctx context.Context, req *role.GetRoleRequest, resp *role.GetRoleResponse) error {
	panic("no implemention")
}

func (r *RoleHandler) InsertRole(ctx context.Context, req *role.InsertRoleRequest, resp *role.InsertRoleResponse) error {
	panic("no implemention")
}

func (r *RoleHandler) DeleteRole(ctx context.Context, req *role.DeleteRoleRequest, resp *role.DeleteRoleResponse) error {
	panic("no implemention")
}

func (r *RoleHandler) UpdateRole(ctx context.Context, req *role.UpdateRoleRequest, resp *role.UpdateRoleResponse) error {
	panic("no implemention")
}
