package cache

import (
	redis "gopkg.in/redis.v5"
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

func New(cli *redis.Client) (ICache, error) {

	return &RedisCache{cli: cli}, nil
}
