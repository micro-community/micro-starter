package db

import (
	"sync"

	"github.com/micro-community/auth/cache"
	"github.com/micro-community/auth/config"
	"github.com/micro-community/auth/db/sql"
	"github.com/micro/go-micro/v3/logger"
	"gorm.io/gorm"
)

var (
	cacheCli      *cache.Client
	db            *gorm.DB
	once          sync.Once
	dbContextType string
)

func init() {
	InitCache()

}

func InitCache() {
	var err error
	cacheCli, err = cache.NewClient(config.Cfg.Redis)
	if err != nil {
		logger.Fatal(err)
	}
}

//BuildDBContext for data
func BuildDBContext(dbCase string) {
	dbContextType = dbCase
	switch dbContextType {
	case "mysql", "sqlite":
		// connet to gorm data source
	case "mongo":
		// connect to mongo
	case "dgraph":
		//connect to dgraph
	default:
		//use memory to mock

	}

}

func DB() *gorm.DB {

	if db != nil {
		return db
	}
	db = sql.NewSQLite(config.Cfg.SQLite)
	once.Do(func() {
		migrate()
	})

	if sqlDB, err := db.DB(); err != nil {
		// SetMaxIdleConns 设置空闲连接池中连接的最大数量
		sqlDB.SetMaxIdleConns(config.Cfg.MaxIdleConns)

		// SetMaxOpenConns 设置打开数据库连接的最大数量。
		sqlDB.SetMaxOpenConns(config.Cfg.MaxOpenConns)

		// SetConnMaxLifetime 设置了连接可复用的最大时间。
		sqlDB.SetConnMaxLifetime(config.Cfg.ConnMaxLifetime)
	}

	return db
}
