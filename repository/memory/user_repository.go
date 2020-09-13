package memory

import (
	"sync"

	"github.com/micro-community/auth/models"
	"github.com/micro-community/auth/repository"
)

type userRepository struct {
	mu    *sync.Mutex
	users []*models.User
}

func NewUserRepository() repository.IUser {
	users := make([]*models.User, 0)
	users = append(users, &models.User{
		Id:       1,
		Name:     "admin",
		Password: "123456",
	})

	return &userRepository{
		mu:    &sync.Mutex{},
		users: users,
	}
}

func (r *userRepository) FindById(id int64) (*models.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, user := range r.users {
		if user.Id == id {
			return user, nil
		}
	}
	return nil, nil
}

func (r *userRepository) FindByName(name string) (*models.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, user := range r.users {
		if user.Name == name {
			return user, nil
		}
	}
	return nil, nil
}

func (r *userRepository) Add(user *models.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	id := int64(len(r.users) + 1)
	user.Id = id

	r.users = append(r.users, user)

	return nil
}

func (r *userRepository) List(page, size int) ([]*models.User, error) {
	return nil, nil
}
