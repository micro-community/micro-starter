package cache

import (
	"context"
	"time"

	"github.com/micro/micro/v3/service/logger"
	"golang.org/x/sync/singleflight"
)

type IClient interface {
	Get(ctx context.Context, key string) ([]byte, error)
	Set(ctx context.Context, key string, v []byte, expire time.Duration) error
}

type SyncClient struct {
	sf     *singleflight.Group
	client IClient
}

func NewSyncClient() *SyncClient {
	return &SyncClient{
		sf: &singleflight.Group{},
	}
}

func (s *SyncClient) Get(ctx context.Context, key string) ([]byte, error) {

	logger.Info("cache sync")
	value, err, _ := s.sf.Do(key, func() (interface{}, error) {
		return s.client.Get(ctx, key)
	})
	return value.([]byte), err
}
