package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

//Event type
type Event struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	UserID    primitive.ObjectID `bson:"_userId,omitempty" json:"_userId"`
	EventType int32              `bson:"eventType,omitempty" json:"eventType"`
	RefID     primitive.ObjectID `bson:"_refId,omitempty" json:"_refId"`
	Status    int32              `bson:"status,omitempty" json:"status"`
	Created   time.Time          `bson:"created,omitempty" json:"created"`
	Updated   time.Time          `bson:"updated,omitempty" json:"updated"`
}



type Log struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	UserID    primitive.ObjectID `bson:"_userId,omitempty" json:"_userId"`
	EventType int32              `bson:"eventType,omitempty" json:"eventType"`
	Created   time.Time          `bson:"created,omitempty" json:"created"`
	Updated   time.Time          `bson:"updated,omitempty" json:"updated"`
}
