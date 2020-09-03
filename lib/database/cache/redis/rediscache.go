package rediscache

import (
	icache "github.com/crazybber/user/lib/database/cache"

	"gopkg.in/redis.v5"
)

type RedisCache struct {
	cli *redis.Client
}

func (r RedisCache) Get() (interface{}, error) {
	panic("implement me")
}

func (r RedisCache) Set() error {
	panic("implement me")
}

func New(cli *redis.Client) (icache.Cache, error) {

	return &RedisCache{cli: cli}, nil
}
