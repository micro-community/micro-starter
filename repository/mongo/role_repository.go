package mongo

import (
	"context"

	"github.com/micro-community/auth/models"
	"github.com/micro/micro/v3/service/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoleRepository struct {
	db *mongo.Collection
}

// ListAllRole - gets all
func (r *RoleRepository) ListAllRole(ctx context.Context) (roles []models.Role, err error) {

	cursor, err := r.db.Find(ctx, bson.D{})
	if err != nil {
		logger.Infof("Find error: %v", err)
		return
	}
	defer cursor.Close(ctx)
	if err = cursor.All(ctx, &roles); err != nil {
		logger.Infof("Error getting data: ", err)
		return
	}

	return
}

// GetByID - gets role by id
func (r *RoleRepository) GetByID(ctx context.Context, ID string) (role models.Role, err error) {

	objID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		logger.Infof("Invalid id: %v", err)
		return
	}
	sr := r.db.FindOne(context.TODO(), bson.M{"_id": objID})
	err = sr.Decode(&role)
	if err != nil {
		logger.Infof("Error decoding data: ", err)
	}

	return
}

// Create - creates new role in collection
func (r *RoleRepository) Create(ctx context.Context, roleName string) (err error) {
	tempRoleID := primitive.NewObjectID()

	if _, err = r.db.InsertOne(ctx, tempRoleID); err != nil {
		logger.Infof("Error inserting role: %v", r)
		return
	}

	logger.Info("Inserted role into collection")
	return
}

// Delete - deletes role
func (r *RoleRepository) Delete(ctx context.Context, roleID string) (err error) {

	if _, err = r.db.DeleteOne(context.TODO(), roleID); err != nil {
		logger.Infof("Error deleting role: %v", roleID)
		return
	}

	logger.Info("Deleted role from collection")
	return
}
