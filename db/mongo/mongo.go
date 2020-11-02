package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/micro/micro/v3/service/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//MongoOptions for mongodb
type MongoOptions struct {
	User     string
	Password string
	Host     string
	Port     int
	DBName   string
}

var (
	db *mongo.Database
)

// Init - init mongodb connection
func Init(config *MongoOptions) {

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", config.User, config.Password, config.Host, config.Port)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		logger.Fatal("Cannot initialize database")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	if err = client.Connect(ctx); err != nil {
		logger.Fatal("Cannot initialize database context")
	}

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		logger.Fatal("Cannot ping database")
	}

	logger.Info("Connected To MongoDB")
	db = client.Database(config.DBName)
	return
}

// GetDB gets db connection
func GetDB() *mongo.Database {
	if db == nil {
		logger.Fatal("Database not initialized")
		return nil
	}

	return db
}
