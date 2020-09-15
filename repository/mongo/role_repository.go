package mongo

import (
	"context"
	"time"

	"github.com/micro-community/auth/models"
	"github.com/micro/go-micro/v3/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


type RoleRepository struct {
	db *mongo.Collection
}

// ListAllRole - gets all
func (r *RoleRepository) ListAllRole() (roles []models.Role, err error) {

	cursor, err := r.db.Find(context.TODO(), bson.D{})
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
func (r *RoleRepository) GetByID(ID string) (role models.Role, err error) {

	objID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		logger.Infof("Invalid id: %v", err)
		return
	}
	sr :=  r.db.FindOne(context.TODO(), bson.M{"_id": objID})
	err = sr.Decode(&role)
	if err != nil {
		logger.Infof("Error decoding data: ", err)
	}

	return
}

// Create - creates new row in collection
func (r *RoleRepository) Create(ctx context.Context) (err error) {
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

	if _, err =  r.db.InsertOne(ctx, r); err != nil {
		logger.Infof("Error inserting role: %v", r)
		return
	}

	logger.Info("Inserted role into collection")
	return
}

// Delete - deletes role
func (r *RoleRepository) Delete() (err error) {

	if _, err =  r.db.DeleteOne(context.TODO(), r); err != nil {
		logger.Infof("Error deleting role: %v", r)
		return
	}

	logger.Info("Deleted role from collection")
	return
}
