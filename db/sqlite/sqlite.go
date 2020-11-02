package sqlite

import (
	_ "gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SQLiteOptions struct {
	User      string
	Password  string
	Host      string
	Port      int
	DBName    string
	Path      string
	LogDetail bool
}

func (cfg *SQLiteOptions) WithSQLiteDefault() *SQLiteOptions {
	if cfg == nil {
		return nil
	}
	if cfg.Port == 0 {
		cfg.Port = 3306
	}
	if cfg.Host == "" {
		cfg.Host = "127.0.0.1"
	}
	return cfg
}

func NewSQLite(cfg *SQLiteOptions) *gorm.DB {

	c := cfg.WithSQLiteDefault()
	// 返回一个连接池
	db, err := gorm.Open(sqlite.Open(c.DBName), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return db
}
