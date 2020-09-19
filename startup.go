package main

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

//serviceCollection for DI.
type serviceCollection struct {
	dig.In
	roleService     *service.RoleService
	userService     *service.UserService
	resourceService *service.ResourceService

	// .... 其他的service
}

//buildingStartupService build all service relationship
func buildingStartupService(srv *mservice.Service) {

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
	db.BuildDBContext(config.Cfg.DefaultDB)
	switch config.Cfg.DefaultDB {
	case "dgraph":
		c.Provide(dgraph.NewRBACRepository)
	default:
		// 默认memory
		c.Provide(memory.NewUserRepository)
	}

}
