package db

import (
	"sync"

	"github.com/micro-community/auth/cache"
	"github.com/micro-community/auth/config"
	"github.com/micro-community/auth/db/nosql"
	"github.com/micro-community/auth/db/sql"
	"github.com/micro/go-micro/v3/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var (
	cacheCli      *cache.Client
	db            *gorm.DB      // for mysql/sqlite
	dg            *nosql.DormDB //for dgraph
	mdb           *mongo.Database
	once          sync.Once
	dbContextType string
)

func InitCache(conf *config.Config) {
	var err error
	cacheCli, err = cache.NewClient(conf.Redis)
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

func DDB() *nosql.DormDB {

	if dg != nil {
		return dg
	}

	dg = nosql.NewDGraphClient(config.Cfg.Dgraph)
	once.Do(func() {
		migrate()
	})

	return dg

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
