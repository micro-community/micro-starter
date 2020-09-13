package repository

import "github.com/micro-community/auth/models"

//IUser for user
type IUser interface {
	FindById(id int64) (*models.User, error)
	FindByName(name string) (*models.User, error)
	Add(user *models.User) error
	List(page, size int) ([]*models.User, error)
}
