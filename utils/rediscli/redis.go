package rediscli

import (
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/micro-community/auth/config"
	"github.com/micro-community/auth/utils"
)

var Redis *redis.Pool

//for cache
var (
	DefaultRedisAddr = "micro.starter.redis:6379"
)

//InitRedisAddr set redis addr
func InitRedisAddr(addr string) {
	if len(addr) > 0 {
		DefaultRedisAddr = addr
	}
}

func init() {

	Redis = &redis.Pool{
		MaxIdle:     config.Cfg.Redis.MaxIdle,
		MaxActive:   config.Cfg.Redis.MaxActive,
		IdleTimeout: time.Duration(config.Cfg.Redis.MaxIdleTimeout) * time.Second,
		Wait:        config.Cfg.Redis.Wait,
		Dial: func() (redis.Conn, error) {
			con, err := redis.DialURL(config.Cfg.Redis.Url)
			if err != nil {
				return nil, err
			}
			return con, nil
		},
	}
	utils.Sugar.Infow("redis inited.")
}

func GetClient() redis.Conn {
	return Redis.Get()
}
