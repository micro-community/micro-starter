package models

import (
	"errors"

	"github.com/crazybber/user/lib/database/global"
)

//RoleModel of system
type RoleModel struct {
	RoleId   int    `json:"roleId" gorm:"primary_key;AUTO_INCREMENT"` // 角色编码
	RoleKey  string `json:"roleKey" gorm:"size:128;"`                 //角色代码
	RoleName string `json:"roleName" gorm:"size:128;"`                // 角色名称
	Admin    bool   `json:"admin" gorm:"size:4;"`                     // 管理员
	RoleIds  []int  `json:"roleIds" gorm:"-"`                         // 对应的角色列表: 单独的role对应一种权限操作
	UpdateBy string `json:"updateBy" gorm:"size:128;"`                //
	//Operations []Operation `json:"operations"`

	Model
}

func (RoleModel) TableName() string {
	return "role"
}

func (role *RoleModel) Get() (RoleModel RoleModel, err error) {
	table := global.DB().Table("sys_role")
	if role.RoleId != 0 {
		table = table.Where("role_id = ?", role.RoleId)
	}
	if role.RoleName != "" {
		table = table.Where("role_name = ?", role.RoleName)
	}
	if err = table.First(&RoleModel).Error; err != nil {
		return
	}

	return
}

func (role *RoleModel) Insert() (id int, err error) {
	i := 0
	global.DB().Table(role.TableName()).Where("role_name=? or role_key = ?", role.RoleName, role.RoleKey).Count(&i)
	if i > 0 {
		return 0, errors.New("角色名称或者角色标识已经存在！")
	}
	role.UpdateBy = ""
	result := global.DB().Table(role.TableName()).Create(&role)
	if result.Error != nil {
		err = result.Error
		return
	}
	id = role.RoleId
	return
}

//Update 修改
func (role *RoleModel) Update(id int) (update RoleModel, err error) {
	if err = global.DB().Table(role.TableName()).First(&update, id).Error; err != nil {
		return
	}

	if role.RoleName != "" && role.RoleName != update.RoleName {
		return update, errors.New("角色名称不允许修改！")
	}

	if role.RoleKey != "" && role.RoleKey != update.RoleKey {
		return update, errors.New("角色标识不允许修改！")
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = global.DB().Table(role.TableName()).Model(&update).Updates(&role).Error; err != nil {
		return
	}
	return
}

//批量删除
func (role *RoleModel) BatchDelete(id []int) (Result bool, err error) {
	if err = global.DB().Table(role.TableName()).Where("role_id in (?)", id).Delete(RoleModel{}).Error; err != nil {
		return
	}
	Result = true
	return
}
