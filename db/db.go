package db

import (
	"sync"

	"github.com/micro-community/micro-starter/cache"
	"github.com/micro-community/micro-starter/config"
	"github.com/micro-community/micro-starter/db/dgraph"
	"github.com/micro-community/micro-starter/db/sqlite"
	"github.com/micro/micro/v3/service/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var (
	cacheCli      *cache.Client
	db            *gorm.DB       // for mysql/sqlite
	dg            *dgraph.DormDB //for dgraph
	mdb           *mongo.Database
	once          sync.Once
	dbContextType string
)

func InitCache(conf *config.Options) {
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

func DDB() *dgraph.DormDB {

	if dg != nil {
		return dg
	}

	dg = dgraph.NewDGraphClient(config.Default.Dgraph)
	once.Do(func() {
		migrate()
	})

	return dg

}

//MDB for mongodb nosql
func MDB() *mongo.Database {

	if mdb != nil {
		return mdb
	}
	//create mong link to connect mongo

	return mdb

}

//DB for mysql/sqlite/pg
func DB() *gorm.DB {

	if db != nil {
		return db
	}
	db = sqlite.NewSQLite(config.Default.SQLite)
	once.Do(func() {
		migrate()
	})

	if sqlDB, err := db.DB(); err != nil {
		// SetMaxIdleConns 设置空闲连接池中连接的最大数量
		sqlDB.SetMaxIdleConns(config.Default.MaxIdleConns)

		// SetMaxOpenConns 设置打开数据库连接的最大数量。
		sqlDB.SetMaxOpenConns(config.Default.MaxOpenConns)

		// SetConnMaxLifetime 设置了连接可复用的最大时间。
		sqlDB.SetConnMaxLifetime(config.Default.ConnMaxLifetime)
	}

	return db
}
