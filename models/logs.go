package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Log struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	UserID    primitive.ObjectID `bson:"_userId,omitempty" json:"_userId"`
	EventType int32              `bson:"eventType,omitempty" json:"eventType"`
	Created   time.Time          `bson:"created,omitempty" json:"created"`
	Updated   time.Time          `bson:"updated,omitempty" json:"updated"`
}
