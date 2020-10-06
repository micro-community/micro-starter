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
		MaxIdle:     config.Default.Redis.MaxIdle,
		MaxActive:   config.Default.Redis.MaxActive,
		IdleTimeout: time.Duration(config.Default.Redis.MaxIdleTimeout) * time.Second,
		Wait:        config.Default.Redis.Wait,
		Dial: func() (redis.Conn, error) {
			con, err := redis.DialURL(config.Default.Redis.Url)
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
