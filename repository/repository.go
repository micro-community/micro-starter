package repository

import "github.com/micro-community/micro-starter/models"

//IUser for user
type IUser interface {
	FindById(id int64) (*models.User, error)
	FindByName(name string) (*models.User, error)
	Add(user *models.User) error
	List(page, size int) ([]*models.User, error)
}

type IRole interface {
	FindById(id int64) (*models.Role, error)
	FindByName(name string) (*models.Role, error)
	Add(user *models.Role) error
	List(page, size int) ([]*models.Role, error)
}

type IResource interface {
	FindById(id int64) (*models.Resource, error)
	FindByName(name string) (*models.Resource, error)
	Add(user *models.Resource) error
	List(page, size int) ([]*models.Resource, error)
}
