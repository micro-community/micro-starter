package memory

import (
	"errors"
	"sync"
	"time"

	"github.com/micro-community/micro-starter/models"
)

//UserModel data
type RoleRepository struct {
	mu    *sync.Mutex
	roles []*models.Role
}

func NewRoleRepository() *RoleRepository {
	roles := make([]*models.Role, 0)
	roles = append(roles, &models.Role{
		ID:   1,
		Name: "boss",
		ModelExtension: models.ModelExtension{
			CreatedAt: time.Now(),
		},
	})

	return &RoleRepository{
		mu:    &sync.Mutex{},
		roles: roles,
	}
}

func (r *RoleRepository) findTarget(roleID int) (int, *models.Role) {

	for index, role := range r.roles {
		if role.ID == roleID {
			return index, r.roles[index]
		}
	}
	return -1, nil
}

func (r *RoleRepository) Get(id int) (*models.Role, error) {

	r.mu.Lock()
	defer r.mu.Unlock()

	_, target := r.findTarget(id)

	return target, nil
}

func (r *RoleRepository) Insert(role *models.Role) (id int, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	id = int(len(r.roles) + 1)
	role.ID = id
	r.roles = append(r.roles, role)
	return
}

//Update 修改
func (r *RoleRepository) Update(update models.Role) (id int, err error) {

	r.mu.Lock()
	defer r.mu.Unlock()

	targetID, targetRole := r.findTarget(update.ID)

	if targetID == -1 {
		return -1, errors.New("target role not exist")
	}

	if update.Key != "" && targetRole.Key != update.Key {
		return -1, errors.New("role key modify forbiden")
	}

	targetRole = &update

	return targetRole.ID, nil
}

func (r *RoleRepository) Del(id int) bool {

	r.mu.Lock()
	defer r.mu.Unlock()

	index, _ := r.findTarget(id)
	if index != -1 {
		r.roles = append(r.roles[:index], r.roles[index+1:]...)
		return true
	}

	return false

}

//BatchDelete 批量删除
func (r *RoleRepository) BatchDelete(ids []int) []bool {

	cnt := len(ids)

	res := make([]bool, cnt)

	for idx, id := range ids {
		res[idx] = r.Del(id)
	}
	return res
}
