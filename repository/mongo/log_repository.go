package mongo

import (
	"context"
	"sync"

	"github.com/micro-community/auth/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type LogRepository struct {
	db *mongo.Collection
	mu *sync.Mutex
}

func (l *LogRepository) Create(ctx context.Context, event models.Log) (interface{}, error) {

	l.mu.Lock()
	defer l.mu.Unlock()

	rsp, err := l.db.InsertOne(ctx, event)
	if err != nil {
		return nil, err
	}
	return rsp.InsertedID, nil
}

func (l *LogRepository) Find(ctx context.Context, query bson.M) ([]models.Log, error) {

	var logs []models.Log
	cursor, err := l.db.Find(ctx, query)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &logs); err != nil {
		return nil, err
	}
	return logs, nil
}

func (l *LogRepository) FindOne(ctx context.Context, query bson.M) (*models.Log, error) {
	var log models.Log
	if err := l.db.FindOne(ctx, query).Decode(&log); err != nil {
		return nil, err
	}
	return &log, nil
}
