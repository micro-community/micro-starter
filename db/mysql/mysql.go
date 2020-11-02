package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySqlOptions struct {
	User      string
	Password  string
	Host      string
	Port      int
	DBName    string
	LogDetail bool

	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime int
}

func (cfg *MySqlOptions) WithMySQLDefault() *MySqlOptions {
	if cfg == nil {
		return nil
	}

	c := *cfg

	if cfg.Port == 0 {
		c.Port = 3306
	}
	if cfg.Host == "" {
		c.Host = "127.0.0.1"
	}

	if cfg.MaxIdleConns == 0 {
		c.MaxIdleConns = 10
	}

	if cfg.MaxOpenConns == 0 {
		c.MaxOpenConns = 80
	}

	if cfg.ConnMaxLifetime == 0 {
		c.ConnMaxLifetime = 0
	}

	return &c
}

// db, err = gorm.Open("mysql", "metro:metro1234@10.252.6.139:3306/crm?charset=utf8mb4")
func (c MySqlOptions) URI() string {
	return fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4", c.User, c.Password, c.Host, c.Port, c.DBName)
}

func NewMySQL(cfg *MySqlOptions) (db *gorm.DB, err error) {

	c := cfg.WithMySQLDefault()
	// 返回一个连接池

	db, err = gorm.Open(mysql.Open(c.URI()), &gorm.Config{})

	// // 防止无线连接，出现 too many connections 错误
	// // MySQL默认是151
	// if c.MaxIdleConns > 0 {
	// 	db.DB().SetMaxIdleConns(c.MaxIdleConns)
	// }
	// if c.MaxOpenConns > 0 {
	// 	db.SetMaxOpenConns(c.MaxOpenConns)
	// }
	// db.SetConnMaxLifetime(time.Duration(c.ConnMaxLifetime))

	// db.LogMode(cfg.LogDetail)

	return
}
