package nosql

import (
	"gopkg.in/redis.v5"
)

type RedisCfg struct {
	MasterName    string   `json:"master_name"`
	SentinelAddrs []string `json:"sentinel_addrs"`

	Host     string `json:"host"`
	Password string `json:"password"`

	DB      int `json:"db"`
	MaxIdle int `json:"max_idle"`
}

type Client struct {
	RedisCfg
	*redis.Client
}

func NewClient(cfg RedisCfg) (client *redis.Client, err error) {

	if len(cfg.SentinelAddrs) != 0 {
		client = redis.NewFailoverClient(&redis.FailoverOptions{
			MasterName:    cfg.MasterName,
			SentinelAddrs: cfg.SentinelAddrs,
			Password:      cfg.Password,

			DB:       cfg.DB,
			PoolSize: cfg.MaxIdle,
		})
	} else {
		client = redis.NewClient(&redis.Options{
			Addr:     cfg.Host,
			Password: cfg.Password,

			DB:       cfg.DB,
			PoolSize: cfg.MaxIdle,
		})
	}

	err = client.Ping().Err()
	return
}
