/*
 * @Author: Edward https://github.com/crazybber
 * @Date: 2020-09-21 17:37:45
 * @Last Modified by: Eamon
 * @Last Modified time: 2020-09-22 23:18:16
 * @Description:  All Service Instance will be created
 */

package profile

import (
	"github.com/micro-community/micro-starter/config"
	"github.com/micro-community/micro-starter/db"
	"github.com/micro-community/micro-starter/handler/rbac"
	"github.com/micro-community/micro-starter/handler/resource"
	"github.com/micro-community/micro-starter/handler/role"
	"github.com/micro-community/micro-starter/handler/user"
	"github.com/micro-community/micro-starter/repository"
	"github.com/micro-community/micro-starter/repository/dgraph"
	"github.com/micro-community/micro-starter/repository/memory"
  "github.com/micro-community/micro-starter/repository/mysql"
	"github.com/micro-community/micro-starter/repository/mongo"
	//	"github.com/micro-community/micro-starter/repository/mysql"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"go.uber.org/dig"
)

//modelService for DI ,all DI　All Service Instance will be created Here
type modelService struct {
	dig.In
	rbac     *dgraph.RbacRepository
	role     *mysql.RoleRepository
	user     repository.IUser
	logs     *mongo.LogRepository

	// .... 其他的service
}

//BuildingStartupService build all service relationship
func BuildingStartupService(srv *service.Service, conf *config.Options) {

	c := dig.New()

	//data(database/cache) context
	buildDataContext(c, conf)

	//service : aggregate repository service and logic proc to provide service ability for handler

	c.Provide(dgraph.NewRBACRepository)
	c.Provide(memory.NewUserRepository)
	c.Provide(mysql.NewRoleRepository)
	c.Provide(mongo.NewLogRepository)
	// begin to handle service object instance
	c.Invoke(func(sc *modelService) {

		if sc == nil {
			logger.Warnf("no service got in DI Container")
		}

		srv.Handle(rbac.NewRBAC(srv, sc.user, sc.role, sc.resource))
		// handle user
		srv.Handle(user.NewUser(srv, sc.user))
		// handle role
		srv.Handle(role.NewRole(srv, sc.role))
		// handle resource
		srv.Handle(resource.NewResource(srv, sc.resource))

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
