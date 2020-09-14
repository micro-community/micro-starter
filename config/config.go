package config

import (
	"github.com/micro-community/auth/cache"
	"github.com/micro-community/auth/db/nosql"
	"github.com/micro-community/auth/db/sql"
	"github.com/micro/go-micro/v3/logger"
	"github.com/micro/micro/v3/service/config"
)

//Service configuration and register
const (
	DbName    = "auth" // database name
	TenantKey = "tenantids"
	BASE_PATH = "./"
)

//Default of config
var Default = &Config{
	Redis: &cache.RedisCfg{
		MasterName:    "",
		SentinelAddrs: nil,
		Host:          "",
		Password:      "",
		DB:            0,
		MaxIdle:       0,
	},
	SQLite: &sql.SQLiteConfig{
		User:     "",
		Password: "",
		Host:     "",
		Port:     0,
		DBName:   "",
		Path:     "",
	},
	Mongodb: &nosql.MongoCfg{
		User:     "",
		Password: "",
		Host:     "",
		Port:     27017,
		DBName:   "",
	},
	Dgraph: &nosql.DgraphCfg{
		User:     "",
		Password: "",
		Host:     "",
		Port:     0,
		DBName:   "",
	},
}

//Config of type
type Config struct {
	Host    string
	Timeout int
	Redis   *cache.RedisCfg
	MySQL   *sql.MySQLConfig
	SQLite  *sql.SQLiteConfig
	Mongodb *nosql.MongoCfg
	Dgraph  *nosql.DgraphCfg
}

var Cfg Config

//读取根目录下的配置，用于初始化配置
func init() {

	//  get config
	svcs := config.Get("micro", "status", "services").StringSlice(nil)
	logger.Infof("Services config %+v", svcs)

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
