package db

import (
	"sync"

	"github.com/go-redis/redis/v8"
	rcache "github.com/micro-community/auth/cache"
	"github.com/micro-community/auth/config"
	"github.com/micro-community/auth/db/sql"
	"github.com/micro/go-micro/v3/logger"
	"gorm.io/gorm"
)

var (
	cfg      *config.Config
	redisCli *redis.Client
	db       *gorm.DB
	cacheCli rcache.ICache
	once     sync.Once
	dbConn   string
)

func init() {
	InitCache()
	BuildDBContext(config.Cfg.DefaultDB)
}

func InitCache() {
	var err error
	redisCli, err = rcache.NewClient(*cfg.Redis)
	if err != nil {
		logger.Fatal(err)
	}
	cacheCli, err = rcache.New(redisCli)
	if err != nil {
		logger.Fatal(err)
	}
}

//BuildDBContext for data
func BuildDBContext(dbCase string) {
	dbConn = dbCase
	switch dbConn {
	case "mysql", "sql":
		// DB初始化
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
	db = sql.NewSQLite(cfg.SQLite)
	once.Do(func() {
		migrate()
	})

	return db
}

// User Model
type User struct {
	gorm.Model
	Name     string
	Code     string
	ID       uint
	Gender   int
	Password string
}

func migrate() {

	//Migrate the schema
	db.AutoMigrate(&User{})

	// Create
	db.Create(&User{Code: "D42", ID: 100})

	// Read
	var user User

	db.First(&user, 1)                 // find product with integer primary key
	db.First(&user, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	db.Model(&user).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&user).Updates(User{ID: 200, Code: "F42"}) // non-zero fields
	db.Model(&user).Updates(map[string]interface{}{"ID": 200, "Code": "F42"})

	// Delete - delete product
	db.Delete(&user, 1)

}
