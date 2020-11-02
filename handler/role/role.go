package role

import (
	"context"

	pb "github.com/micro-community/micro-starter/protos"
	"github.com/micro-community/micro-starter/repository"
	"github.com/micro/micro/v3/service"
)

//RoleHandler implements the role proto interface
type RoleHandler struct {
	RoleID string
	Name   string
	role   repository.IRole
}

// NewRole returns an initUser handler
func NewRole(service *service.Service, roleService repository.IRole) *RoleHandler {
	return &RoleHandler{
		Name: service.Name(),
		role: roleService,
	}
}

func (r *RoleHandler) GetRole(ctx context.Context, req *pb.GetRoleRequest, resp *pb.GetRoleResponse) error {
	panic("no implemention")
}

func (r *RoleHandler) InsertRole(ctx context.Context, req *pb.InsertRoleRequest, resp *pb.InsertRoleResponse) error {
	panic("no implemention")
}

func (r *RoleHandler) DeleteRole(ctx context.Context, req *pb.DeleteRoleRequest, resp *pb.DeleteRoleResponse) error {
	panic("no implemention")
}

func (r *RoleHandler) UpdateRole(ctx context.Context, req *pb.UpdateRoleRequest, resp *pb.UpdateRoleResponse) error {
	panic("no implemention")
}
