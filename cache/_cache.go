package cache

import (
	"reflect"
	"time"

	"google.golang.org/protobuf/proto"
	"gopkg.in/redis.v5"

	"golang.org/x/sync/singleflight"
)

type cacheClient int

const (
	Redis cacheClient = iota
	Memcache
)

type Cache struct {
	sf     *singleflight.Group
	client Client
}

func New(client cacheClient) *Cache {
	return &Cache{
		sf: &singleflight.Group{},
	}
}

func (c *Cache) FetchWithProtobuf(key string, expire time.Duration, fn func() (interface{}, error)) error {
	//v, err, _ := c.sf.Do(key, f)

	dl := dlog.FromContext(ctx)
	// 需要穿透cache，取source数据,并覆盖cache
	if cachectx.FromContext(ctx) {
		dl.Info("force fetch source")
		tmp, err := fn()
		if err != nil {
			return nil, err
		}
		if bb, err := proto.Marshal(tmp); err != nil {
			dl.Error("json.Marshal failed", err)
			return nil, err
		} else {
			if err := rediscli.Set(redisKey, bb, redisExpire).Err(); err != nil {
				dl.Errorf("redis.Ser(%v) failed: %v", redisKey, err)
			}
		}
		return tmp, nil
	}
	val, err := rediscli.Get(redisKey).Bytes()
	if err != nil {
		if err == redis.Nil {
			tmp, err := fn()
			var ret = reflect.New(typ)
			if err != nil {
				return nil, err
			}
			if bb, err := proto.Marshal(tmp); err != nil {
				dl.Error("json.Marshal failed", err)
				return nil, err
			} else {
				if err := rediscli.Set(redisKey, bb, redisExpire).Err(); err != nil {
					dl.Errorf("redis.Ser(%v) failed: %v", redisKey, err)
				}
				if err = proto.Unmarshal(bb, ret.Interface().(proto.Message)); err != nil {
					return nil, err
				}
			}
			return ret.Interface().(proto.Message), nil
		} else {
			dl.Errorf("rediscli.Get(%v) failed: %v", redisKey, err)
			return nil, err
		}
	} else {
		var tmp = reflect.New(typ)
		err := proto.Unmarshal(val, tmp.Interface().(proto.Message))
		if err != nil {
			dl.Error("fail to json.Unmarshal", err)
			return nil, err
		}
		return tmp.Interface().(proto.Message), nil
	}
}
