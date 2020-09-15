package mongo

import (
	"context"
	"sync"

	"github.com/micro-community/auth/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//EventRepository for event store
type EventRepository struct {
	db *mongo.Collection
	mu *sync.Mutex
}

func init() {
	//EventDao = &EventRepository{db:mongodb.Client.Collection("event")}
}

func (e *EventRepository) Create(ctx context.Context, event models.Event) (interface{}, error) {

	e.mu.Lock()
	defer e.mu.Unlock()

	rsp, err := e.db.InsertOne(ctx, event)
	if err != nil {
		return nil, err
	}
	return rsp.InsertedID, nil
}

func (e *EventRepository) Find(ctx context.Context, query bson.M) ([]models.Event, error) {
	var events []models.Event
	cursor, err := e.db.Find(ctx, query)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(ctx, &events); err != nil {
		return nil, err
	}
	return events, nil
}

func (e *EventRepository) FindOne(ctx context.Context, query bson.D) (*models.Event, error) {
	var event models.Event
	if err := e.db.FindOne(ctx, query).Decode(&event); err != nil {
		return nil, err
	}
	return &event, nil
}

func (e *EventRepository) UpdateOne(ctx context.Context, query, update bson.D) (*mongo.UpdateResult, error) {

	e.mu.Lock()
	defer e.mu.Unlock()

	rsp, err := e.db.UpdateOne(ctx, query, update)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

func (e *EventRepository) Delete(ctx context.Context, event models.Event) (err error) {

	if _, err = e.db.DeleteOne(context.TODO(), e); err != nil {
		return
	}
	return
}
