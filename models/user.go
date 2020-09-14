package models

import (
	"time"
)

//User Models for db
type User struct {
	ID       int64  `gorm:"primary_key;AUTO_INCREMENT"  json:"id"` // 编码
	Name     string `gorm:"size:64" json:"name"`
	Password string `gorm:"size:128" json:"password"`
	Roles    []int  `json:"roleIds" gorm:"-"` // 对应的角色列表: 单独的role对应一种权限操作
	UserDetails
	UserExtension
}

type UserExtension struct {
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
	IsSoftDel bool      `json:"isSoftDelete"`
}

type UserDetails struct {
	NickName string `gorm:"size:128" json:"nickName"`                                       // 昵称
	Phone    string `gorm:"size:11" json:"phone"`                                           // 手机号
	RoleId   int    `gorm:"" json:"roleId"`                                                 // 角色编码
	DeptId   int    `gorm:"" json:"deptId"`                                                 //部门编码
	PostId   int    `gorm:"" json:"postId"`                                                 //职位编码
	Avatar   string `gorm:"size:255" json:"avatar"`                                         //头像
	Gender   int    `gorm:"type:enum('published', 'pending', 'deleted');default:'pending'"` //性别
	Email    string `gorm:"size:128" json:"email"`                                          //邮箱
}
