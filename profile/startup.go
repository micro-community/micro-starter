/*
 * @Author: Edward https://github.com/crazybber
 * @Date: 2020-09-21 17:37:45
 * @Last Modified by: Eamon
 * @Last Modified time: 2020-09-22 23:18:16
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
	mservice "github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
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
func BuildingStartupService(srv *mservice.Service, conf *config.Options) {

	c := dig.New()

	//data(database/cache) context
	buildDataContext(c, conf)

	//service : aggregate repository service and logic proc to provide service ability for handler
	c.Provide(service.NewUser)
	c.Provide(service.NewRole)
	c.Provide(service.NewResource)

	// begin to handle service object instance
	c.Invoke(func(sc *serviceCollection) {

		if sc == nil {
			logger.Warnf("no service got in DI Container")
		}

		srv.Handle(handler.NewRBAC(srv, sc.userService, sc.roleService, sc.resourceService))
		// handle user
		srv.Handle(handler.NewUser(srv, sc.userService))
		// handle role
		srv.Handle(handler.NewRole(srv, sc.roleService))
		// handle resource
		srv.Handle(handler.NewResource(srv, sc.resourceService))

	})

}

func buildDataContext(c *dig.Container, conf *config.Options) {

	db.BuildDBContext(conf.DBType)

	switch conf.DBType {
	case "dgraph":
		c.Provide(dgraph.NewRBACRepository)
	default:
		// 默认memory
		c.Provide(memory.NewUserRepository)
	}

	db.InitCache(conf)

}
