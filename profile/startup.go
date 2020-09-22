/*
 * @Author: Edward https://github.com/crazybber
 * @Date: 2020-09-21 17:37:45
 * @Last Modified by: Eamon
 * @Last Modified time: 2020-09-21 17:38:29
 * @Description:  All Service Instance will be created
 */

package profile

import (
	"github.com/micro-community/auth/config"
	"github.com/micro-community/auth/db"
	"github.com/micro-community/auth/handler"
	"github.com/micro-community/auth/repository/dgraph"
	"github.com/micro-community/auth/repository/memory"
	"github.com/micro-community/auth/service"
	"github.com/micro/go-micro/v3/logger"
	mservice "github.com/micro/micro/v3/service"
	"go.uber.org/dig"
)

//serviceCollection for DI ,all DI　All Service Instance will be created Here
type serviceCollection struct {
	dig.In
	roleService     *service.RoleService
	userService     *service.UserService
	resourceService *service.ResourceService

	// .... 其他的service
}

//BuildingStartupService build all service relationship
func BuildingStartupService(srv *mservice.Service) {

	c := dig.New()

	//db context
	buildDBContext(c)

	//service : aggregate repository service and logic proc to provide service ability for handler
	c.Provide(service.NewUser)
	c.Provide(service.NewRole)
	c.Provide(service.NewResource)

	// begin to handle service object instance
	c.Invoke(func(sc *serviceCollection) {

		if sc == nil {
			logger.Warnf("no service got in DI Container")
		}
		// handle user
		srv.Handle(handler.NewUser(srv, sc.userService))
		// handle role
		srv.Handle(handler.NewRole(srv, sc.roleService))
		// handle resource
		srv.Handle(handler.NewResource(srv, sc.resourceService))

	})

}

func buildDBContext(c *dig.Container) {
	db.BuildDBContext(config.Cfg.DBType)
	switch config.Cfg.DBType {
	case "dgraph":
		c.Provide(dgraph.NewRBACRepository)
	default:
		// 默认memory
		c.Provide(memory.NewUserRepository)
	}

	db.InitCache()

}
