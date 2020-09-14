package mysql

import (
	"errors"
	"sync"

	"github.com/micro-community/auth/db"
	"github.com/micro-community/auth/models"
)

//UserModel data
type RoleRepository struct {
	tableName string
	mu        *sync.Mutex
	roles     []*models.Role
	
}

func (r RoleRepository) TableName() string {
	return "role"
}

func (r *RoleRepository) Get() (role models.Role, err error) {

	table := db.DB().Table("sys_role")
	if r.ID != 0 {
		table = table.Where("role_id = ?", r.ID)
	}
	if r.Name != "" {
		table = table.Where("role_name = ?", r.Name)
	}
	if err = table.First(&role).Error; err != nil {
		return
	}

	return
}

func (role *RoleRepository) Insert() (id int, err error) {
	var i int64 = 0
	db.DB().Table(role.TableName()).Where("role_name=? or role_key = ?", role.Name, role.Key).Count(&i)
	if i > 0 {
		return 0, errors.New("角色名称或者角色标识已经存在！")
	}
	role.UpdateBy = ""
	result := db.DB().Table(role.TableName()).Create(&role)
	if result.Error != nil {
		err = result.Error
		return
	}
	id = role.ID
	return
}

//Update 修改
func (role *RoleRepository) Update(id int) (update models.Role, err error) {
	if err = db.DB().Table(role.TableName()).First(&update, id).Error; err != nil {
		return
	}

	if role.Name != "" && role.Name != update.Name {
		return update, errors.New("角色名称不允许修改！")
	}

	if role.Key != "" && role.Key != update.Key {
		return update, errors.New("角色标识不允许修改！")
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = db.DB().Table(role.TableName()).Model(&update).Updates(&role).Error; err != nil {
		return
	}
	return
}

//批量删除
func (role *RoleRepository) BatchDelete(id []int) (Result bool, err error) {
	if err = db.DB().Table(role.TableName()).Where("role_id in (?)", id).Delete(Role{}).Error; err != nil {
		return
	}
	Result = true
	return
}
