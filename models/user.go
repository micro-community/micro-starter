package models

import (
	"errors"
	"time"

	"github.com/crazybber/user/lib/database/global"
	user "github.com/crazybber/user/proto"

	"golang.org/x/crypto/bcrypt"
)

type Model struct {
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

type UserName struct {
	Username string `gorm:"size:64" json:"username"`
}

type PassWord struct {
	// 密码
	Password string `gorm:"size:128" json:"password"`
}

type LoginM struct {
	UserName
	PassWord
}

type UserIdM struct {
	UserId int64 `gorm:"primary_key;AUTO_INCREMENT"  json:"userId"` // 编码
}

type UserBase struct {
	NickName string `gorm:"size:128" json:"nickName"`                                       // 昵称
	Phone    string `gorm:"size:11" json:"phone"`                                           // 手机号
	RoleId   int    `gorm:"" json:"roleId"`                                                 // 角色编码
	DeptId   int    `gorm:"" json:"deptId"`                                                 //部门编码
	PostId   int    `gorm:"" json:"postId"`                                                 //职位编码
	Avatar   string `gorm:"size:255" json:"avatar"`                                         //头像
	Sex      int    `gorm:"type:enum('published', 'pending', 'deleted');default:'pending'"` //性别
	Email    string `gorm:"size:128" json:"email"`                                          //邮箱

	Model
}

type UserModel struct {
	UserIdM
	UserBase
	LoginM
}

func (UserModel) TableName() string {
	return "user"
}

// 获取用户数据
func (u UserModel) Get() (UserModel UserModel, err error) {
	table := global.DB().Table(u.TableName()).Select([]string{"user.*", "role.role_name"})
	table = table.Joins("left join role on user.role_id=role.role_id")
	if u.UserId != 0 {
		table = table.Where("user_id = ?", u.UserId)
	}

	if u.Username != "" {
		table = table.Where("username = ?", u.Username)
	}

	if u.Password != "" {
		table = table.Where("password = ?", u.Password)
	}

	if u.RoleId != 0 {
		table = table.Where("role_id = ?", u.RoleId)
	}

	if u.DeptId != 0 {
		table = table.Where("dept_id = ?", u.DeptId)
	}

	if u.PostId != 0 {
		table = table.Where("post_id = ?", u.PostId)
	}

	if err = table.First(&UserModel).Error; err != nil {
		return
	}

	UserModel.Password = ""
	return
}

//加密
func (u *UserModel) Encrypt() (err error) {
	if u.Password == "" {
		return
	}

	var hash []byte
	if hash, err = bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost); err != nil {
		return
	} else {
		u.Password = string(hash)
		return
	}
}

//添加
func (u UserModel) Insert() (id int64, err error) {
	if err = u.Encrypt(); err != nil {
		return
	}

	// check 用户名
	var count int
	global.DB().Table(u.TableName()).Where("username = ?", u.Username).Count(&count)
	if count > 0 {
		err = errors.New("账户已存在！")
		return
	}

	//添加数据
	if err = global.DB().Table(u.TableName()).Create(&u).Error; err != nil {
		return
	}
	id = u.UserId
	return
}

//修改
func (u *UserModel) Update(id int64) (update UserModel, err error) {
	if u.Password != "" {
		if err = u.Encrypt(); err != nil {
			return
		}
	}
	if err = global.DB().Table(u.TableName()).First(&update, id).Error; err != nil {
		return
	}
	if u.RoleId == 0 {
		u.RoleId = update.RoleId
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = global.DB().Table(u.TableName()).Model(&update).Updates(&u).Error; err != nil {
		return
	}
	return
}
func (u *UserModel) BatchDelete(id []int) (Result bool, err error) {
	if err = global.DB().Table(u.TableName()).Where("user_id in (?)", id).Delete(&UserModel{}).Error; err != nil {
		return
	}
	Result = true
	return
}

func (u *UserModel) ToView() *user.UserInfo {
	var v user.UserInfo
	v.Name = u.NickName
	//.....

	return &v
}
