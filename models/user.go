package models

import (
	"time"
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
