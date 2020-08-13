package cache

import "time"

type Client interface {
	Get(key string, v string, expire time.Duration, f func() (interface{}, error)) (interface{}, error)
	Set() error
}
