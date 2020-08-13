package global

import (
	"auth-demo/config"
	icache "auth-demo/lib/database/cache"
	rediscache "auth-demo/lib/database/cache/redis"
	"auth-demo/lib/database/nosql"

	"github.com/micro/go-micro/v3/logger"

	"gopkg.in/redis.v5"

	"github.com/jinzhu/gorm"
)

var (
	cfg      *config.Config
	redisCli *redis.Client
	cache    icache.Cache

	db *gorm.DB
)

func Init() {
	//github.com/micro/go-plugins/logger/zap
	//logger.DefaultLogger = zap.NewLogger()

	var err error
	cfg = config.Load(nil)

	redisCli, err = nosql.NewClient(*cfg.Redis)
	if err != nil {
		logger.Fatal(err)
	}

	cache, err = rediscache.New(redisCli)
	if err != nil {
		logger.Fatal(err)
	}
}

func DB() *gorm.DB {
	if db == nil {
		return nil
	}

	return db.New()
}
