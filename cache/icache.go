package cache

import "time"

type ICache interface {
	Get() (interface{}, error)
	Set() error
}

type IClient interface {
	Get(key string, v string, expire time.Duration, f func() (interface{}, error)) (interface{}, error)
	Set() error
}
