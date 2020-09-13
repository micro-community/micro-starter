package redis

import (
	"time"

	redigo "github.com/gomodule/redigo/redis"
	"github.com/micro-community/auth/config"
	"github.com/micro-community/auth/utils"
)

var Redis *redigo.Pool

//for cache
var (
	RedisAddr = "micro.srv.redis:6379"
)

//InitRedisAddr set redis addr
func InitRedisAddr(addr string) {
	if len(addr) > 0 {
		RedisAddr = addr
	}
}

func init() {

	Redis = &redigo.Pool{
		MaxIdle:     config.Cfg.Redis.MaxIdle,
		MaxActive:   config.Cfg.Redis.MaxActive,
		IdleTimeout: time.Duration(config.Cfg.Redis.MaxIdleTimeout) * time.Second,
		Wait:        config.Cfg.Redis.Wait,
		Dial: func() (redigo.Conn, error) {
			con, err := redigo.DialURL(config.Cfg.Redis.Url)
			if err != nil {
				return nil, err
			}
			return con, nil
		},
	}
	utils.Sugar.Infow("redis inited.")
}

func GetClient() redigo.Conn {
	return Redis.Get()
}
