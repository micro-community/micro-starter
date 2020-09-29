/*
 * @Description: https://github.com/crazybber
 * @Author: Edward
 * @Date: 2020-09-16 23:14:15
 * @Last Modified by: Eamon
 * @Last Modified time: 2020-09-16 23:35:01
 */

package cache

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/micro/go-micro/v3/logger"
)

//Options for redis cache
type Options struct {
	MasterName    string   `json:"master_name"`
	SentinelAddrs []string `json:"sentinel_addrs"`

	Host     string `json:"host"`
	Password string `json:"password"`

	DB             int    `json:"db"`
	MaxIdle        int    `json:"max_idle"`
	Url            string `json:"url"`
	MaxActive      int    `json:"max_active"`
	Wait           bool   `json:"wait"`
	MaxIdleTimeout int    `json:"max_idle_timeout"`
}

type Client struct {
	cfg Options
	cli *redis.Client
}

var (
	cli    *Client
	config Options
	once   sync.Once
)

func (c *Client) Get(ctx context.Context, key string) ([]byte, error) {
	value, err := c.cli.Get(ctx, key).Bytes()
	if err != nil && err == redis.Nil {
		return nil, errors.New("key not exist")
	}
	return value, nil
}

func (c *Client) Set(ctx context.Context, key string, v []byte, expire time.Duration) error {

	if err := c.cli.Set(ctx, key, v, expire).Err(); err != nil {
		logger.Errorf("redis.Set(%v) failed: %v", key, err)
		return err
	}
	return nil
}

func NewClient(cfg Options) (client *Client, err error) {

	if cli != nil {
		return cli, nil
	}

	var redisClient *redis.Client

	//for sentinel HA cluster
	if len(cfg.SentinelAddrs) != 0 {
		redisClient = redis.NewFailoverClient(&redis.FailoverOptions{
			MasterName:    cfg.MasterName,
			SentinelAddrs: cfg.SentinelAddrs,
			Password:      cfg.Password,
			DB:            cfg.DB,
			PoolSize:      cfg.MaxIdle,
		})

	} else {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     cfg.Host,
			Password: cfg.Password,
			DB:       cfg.DB,
			PoolSize: cfg.MaxIdle,
		})

	}

	client = &Client{
		cli: redisClient,
		cfg: cfg,
	}

	cli = client
	config = cfg

	err = redisClient.Ping(context.TODO()).Err()
	return
}
