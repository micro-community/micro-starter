package config

import (
	"github.com/crazybber/user/lib/database/sql"

	"github.com/crazybber/user/lib/database/nosql"

	"github.com/micro/go-micro/v3/logger"
)

//Default of config
var Default = &Config{
	Redis: &nosql.RedisCfg{
		MasterName:    "",
		SentinelAddrs: nil,
		Host:          "",
		Password:      "",
		DB:            0,
		MaxIdle:       0,
	},
	MySql: &sql.MySQLConfig{
		User:            "",
		Password:        "",
		Host:            "",
		Port:            0,
		DBName:          "",
		MaxIdleConns:    0,
		MaxOpenConns:    0,
		ConnMaxLifetime: 0,
	},
}

//Config of type
type Config struct {
	Redis *nosql.RedisCfg
	MySql *sql.MySQLConfig
}

func Load(fn func() (*Config, error)) *Config {
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
