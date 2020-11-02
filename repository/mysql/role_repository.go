package mysql

import (
	"errors"
	"sync"
	"time"

	"github.com/micro-community/micro-starter/models"
	"gorm.io/gorm"
)

//RoleRepository data
type RoleRepository struct {
	mu        *sync.Mutex
	db        *gorm.DB
	tableName string
}

func NewRoleRepository() *RoleRepository {

	return &RoleRepository{}

}

func (r RoleRepository) TableName() string {
	if r.tableName == "" {
		r.tableName = "role"
	}
	return r.tableName
}
func (r RoleRepository) Table() *gorm.DB {
	return r.Table()
}

func (r *RoleRepository) Get(role models.Role) (models.Role, error) {

	table := r.Table()
	if role.ID != 0 {
		table = table.Where("role_id = ?", role.ID)
	}
	if role.Name != "" {
		table = table.Where("role_name = ?", role.Name)
	}
	var result models.Role
	err := table.First(&result).Error

	if err != nil {
		return models.Role{}, err
	}

	return result, nil
}

func (r *RoleRepository) Insert(role models.Role) (id int, err error) {
	var i int64 = 0
	r.Table().Where("role_name=? or role_key = ?", role.Name, role.Key).Count(&i)
	if i > 0 {
		return 0, errors.New("role exist")
	}

	role.CreatedAt = time.Now()
	result := r.Table().Create(&role)
	if result.Error != nil {
		err = result.Error
		return
	}
	id = role.ID
	return
}

//Update 修改
func (r *RoleRepository) Update(update models.Role) (id int, err error) {

	var targetRole models.Role

	if err = r.Table().First(&targetRole, update.ID).Error; err != nil {
		return -1, errors.New("target role not exist")
	}

	if update.Key != "" && targetRole.Key != update.Key {
		return -1, errors.New("role key modify forbiden")
	}

	if err = r.Table().Model(&targetRole).Updates(&update).Error; err != nil {
		return -1, errors.New("update role error")
	}

	return targetRole.ID, nil
}

//批量删除
func (r *RoleRepository) BatchDelete(id []int) (Result bool, err error) {
	if err = r.Table().Where("role_id in (?)", id).Delete(models.Role{}).Error; err != nil {
		return
	}
	Result = true
	return
}
