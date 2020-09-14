package mongo

import (
	"context"
	"time"

	"github.com/micro/go-micro/v3/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// RoleTable - collection name
const RoleTable string = "role"

// Role - role
type Role struct {
	db *mongo.Database
	// Database specific fields
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	UserID    primitive.ObjectID `bson:"user_id,omitempty" json:"user_id,omitempty"`
	CreatedAt time.Time          `bson:"created_at"`
}

// ListAllRole - gets all
func (r *Role) ListAllRole() (roles []Role, err error) {

	collection := r.db.Collection(RoleTable)
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		logger.Infof("Find error: %v", err)
		return
	}
	defer cursor.Close(context.TODO())
	if err = cursor.All(context.TODO(), &roles); err != nil {
		logger.Infof("Error getting data: ", err)
		return
	}

	return
}

// GetByID - gets role by id
func (r *Role) GetByID(ID string) (role Role, err error) {

	collection := r.db.Collection(RoleTable)
	objID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		logger.Infof("Invalid id: %v", err)
		return
	}

	sr := collection.FindOne(context.TODO(), bson.M{"_id": objID})
	err = sr.Decode(&role)
	if err != nil {
		logger.Infof("Error decoding data: ", err)
	}

	return
}

// Create - creates new row in collection
func (r *Role) Create(ctx context.Context) (err error) {
	var tempUserID primitive.ObjectID
	userID := ""

	if ctx.Value("UserID") != nil {
		userID = ctx.Value("UserID").(string)
	}

	if userID == "" {
		logger.Error("User not specified in context")
		tempUserID = primitive.NewObjectID()
	}
	r.UserID = tempUserID
	r.CreatedAt = time.Now()

	collection := r.db.Collection(RoleTable)
	if _, err = collection.InsertOne(ctx, r); err != nil {
		logger.Infof("Error inserting role: %v", r)
		return
	}

	logger.Info("Inserted role into collection")
	return
}

// Delete - deletes role
func (r *Role) Delete() (err error) {

	collection := r.db.Collection(RoleTable)

	if _, err = collection.DeleteOne(context.TODO(), r); err != nil {
		logger.Infof("Error deleting role: %v", r)
		return
	}

	logger.Info("Deleted role from collection")
	return
}
