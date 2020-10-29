/*
 * @Author: Edward https://github.com/crazybber
 * @Date: 2020-09-22 22:28:55
 * @Last Modified by: Eamon
 * @Last Modified time: 2020-09-22 22:39:47
 * @Description: Current file for graph database(dgraph) using
 */

package rbac

import (
	"context"

	rbac "github.com/micro-community/micro-starter/protos/rbac"
	"github.com/micro-community/micro-starter/service"
	mService "github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

//RbacHandler implements the user proto interface,only for rbac

type RbacHandler struct {
	Name        string
	UserSrv     *service.UserService     // instance of the user service
	RoleSrv     *service.RoleService     // instance of the role service
	ResourceSrv *service.ResourceService // instance of the resource service
}

func NewRBAC(service *mService.Service,
	user *service.UserService,
	role *service.RoleService,
	resource *service.ResourceService) *RbacHandler {
	return &RbacHandler{
		Name:        service.Name(),
		UserSrv:     user,
		RoleSrv:     role,
		ResourceSrv: resource,
	}
}

// AddUser is a single request handler called via client.AddUser or the generated client code
func (r *RbacHandler) AddUser(ctx context.Context, req *rbac.User, rsp *rbac.Response) error {
	logger.Infof("Received RbacHandler.AddUser request, ID: %s, Name: %s", req.Id, req.Name)

	//	rsp.Msg = fmt.Sprintf("person created, id: %s,  uid: %s", req.Id, result.Uids[req.Id])
	return nil
}

// RemoveUser is a single request handler called via client.RemoveUser or the generated client code
func (r *RbacHandler) RemoveUser(ctx context.Context, req *rbac.Request, rsp *rbac.Response) error {
	logger.Infof("Received RbacHandler.RemoveUser request, ID: %s", req.Id)

	rsp.Msg = "OK"
	return nil
}

// QueryUserRoles is a single request handler called via client.QueryUserRoles or the generated client code
func (r *RbacHandler) QueryUserRoles(ctx context.Context, req *rbac.Request, rsp *rbac.Roles) error {
	logger.Infof("Received RbacHandler.QueryUserRoles request, ID: %s", req.Id)

	return nil
}

// QueryUserResources is a single request handler called via client.QueryUserResources or the generated client code
func (r *RbacHandler) QueryUserResources(ctx context.Context, req *rbac.Request, rsp *rbac.Resources) error {
	logger.Infof("Received RbacHandler.QueryUserResources request, ID: %s", req.Id)

	return nil
}

// LinkUserRole is a single request handler called via client.LinkUserRole or the generated client code
func (r *RbacHandler) LinkUserRole(ctx context.Context, req *rbac.LinkRequest, rsp *rbac.Response) error {
	logger.Info("Received RbacHandler.LinkUserRole(Add a role for user) request: id1: %s, id2: %s", req.Id1, req.Id2)

	return nil
}

// UnlinkUserRole is a single request handler called via client.UnlinkUserRole or the generated client code
func (r *RbacHandler) UnlinkUserRole(ctx context.Context, req *rbac.LinkRequest, rsp *rbac.Response) error {
	logger.Infof("Received RbacHandler.UnlinkUserRole(Remove a role from user) request: id1: %s, id2: %s", req.Id1, req.Id2)

	return nil
}

// AddRole is a single request handler called via client.AddRole or the generated client code
func (r *RbacHandler) AddRole(ctx context.Context, req *rbac.Role, rsp *rbac.Response) error {
	logger.Infof("Received RbacHandler.AddRole request, ID: %s, Name: %s", req.Id, req.Name)

	return nil
}

// RemoveRole is a single request handler called via client.RemoveRole or the generated client code
func (r *RbacHandler) RemoveRole(ctx context.Context, req *rbac.Request, rsp *rbac.Response) error {
	logger.Infof("Received RbacHandler.RemoveRole request, ID: %s", req.Id)

	return nil
}

// QueryRoleResources is a single request handler called via client.QueryRoleResources or the generated client code
func (r *RbacHandler) QueryRoleResources(ctx context.Context, req *rbac.Request, rsp *rbac.Resources) error {
	logger.Infof("Received RbacHandler.QueryRoleResources request, ID: %s", req.Id)

	return nil
}

// LinkRoleResource is a single request handler called via client.LinkRoleResource or the generated client code
func (r *RbacHandler) LinkRoleResource(ctx context.Context, req *rbac.LinkRequest, rsp *rbac.Response) error {
	logger.Info("Received RbacHandler.LinkRoleResource request: id1: %s, id2: %s", req.Id1, req.Id2)

	return nil
}

// UnlinkRoleResource is a single request handler called via client.UnlinkRoleResource or the generated client code
func (r *RbacHandler) UnlinkRoleResource(ctx context.Context, req *rbac.LinkRequest, rsp *rbac.Response) error {
	logger.Info("Received RbacHandler.UnlinkRoleResource request: id1: %s, id2: %s", req.Id1, req.Id2)

	return nil
}

// AddResource is a single request handler called via client.AddResource or the generated client code
func (r *RbacHandler) AddResource(ctx context.Context, req *rbac.Resource, rsp *rbac.Response) error {
	logger.Infof("Received RbacHandler.AddResource request, ID: %s, Name: %s", req.Id, req.Name)

	return nil
}

// RemoveResource is a single request handler called via client.RemoveResource or the generated client code
func (r *RbacHandler) RemoveResource(ctx context.Context, req *rbac.Request, rsp *rbac.Response) error {
	logger.Infof("Received RbacHandler.RemoveResource request, ID: %s", req.Id)

	rsp.Msg = "OK"
	return nil
}
