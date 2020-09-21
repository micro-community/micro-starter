/*
 * @Author: Edward https://github.com/crazybber
 * @Date: 2020-09-21 17:43:58
 * @Last Modified by: Eamon
 * @Last Modified time: 2020-09-21 19:12:04
 * @Description: Current File for Configuration of current service
 */

package config

import (
	"time"

	"github.com/micro-community/auth/cache"
	"github.com/micro-community/auth/db/nosql"
	"github.com/micro-community/auth/db/sql"
	"github.com/micro/go-micro/v3/logger"
	"github.com/micro/micro/v3/service/config"
)

//读取根目录下的配置，用于初始化配置
func init() {
	//  get config
	svcs := config.Get("micro", "status", "services").StringSlice(nil)
	logger.Infof("Services config %+v", svcs)

}

//Config of type
type Config struct {
	DefaultDB string
	Host      string
	Timeout   int
	Redis     cache.RedisCfg
	MySQL     *sql.MySQLConfig
	SQLite    *sql.SQLiteConfig
	Mongodb   *nosql.MongoCfg
	Dgraph    *nosql.DgraphCfg

	// sql db config
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

//Service configuration and register
var (
	TenantKey      = "tenantids"
	BASE_HERF_PATH = "./"
	Cfg            Config //User Loaded COnfig,if not setted ,default value will be used.
)

//Default of config
var Default = &Config{

	DefaultDB: "memory",

	MaxOpenConns:    10,
	MaxIdleConns:    100,
	ConnMaxLifetime: time.Duration(time.Hour),

	Redis: cache.RedisCfg{
		MasterName:     "",
		SentinelAddrs:  nil,
		Host:           "localhost",
		Password:       "",
		DB:             0,
		MaxIdle:        1,
		MaxIdleTimeout: 1,
	},
	SQLite: &sql.SQLiteConfig{
		User:     "",
		Password: "",
		Host:     "localhost",
		Port:     0,
		DBName:   "",
		Path:     "",
	},
	Mongodb: &nosql.MongoCfg{
		User:     "",
		Password: "",
		Host:     "localhost",
		Port:     27017,
		DBName:   "",
	},
	Dgraph: &nosql.DgraphCfg{
		User:     "",
		Password: "",
		Host:     "localhost",
		Port:     0,
		DBName:   "",
	},
}

//LoadConfigWithDefault Load Config With Default
func LoadConfigWithDefault(fn func() (*Config, error)) *Config {

	if fn == nil {
		logger.Warnf("use default config")
		return Default
	}
	cfg, err := fn()

	if err != nil {
		logger.Warnf("load config failed: %v, use default", err)
		return Default
	}
	return cfg
}
