package models

type UID struct {
	UID string `json:"uid"`
}

//User Models for db
type User struct {
	Uid      string `gorm:"-" json:"uid,omitempty"`
	ID       int64  `gorm:"primary_key;AUTO_INCREMENT" json:"id"` // 编码
	Type     string `gorm:"size:8" json:"dgraph.type,omitempty"`
	NickName string `gorm:"size:64" json:"nickName"` // 昵称
	Name     string `gorm:"size:64" json:"name"`
	Age      int64  `gorm:"size:3" json:"age,omitempty"`
	Gender   string `gorm:"type:enum('0', '1', '2');default:'0'" json:"gender,omitempty"`
	Password string `gorm:"size:128" json:"password"`
	Key      string `gorm:"size:128" json:"key"`
	Roles    []int  `gorm:"-" json:"roles,omitempty"` // 对应的角色列表: 单独的role对应一种权限操作
	UserDetails
	ModelExtension
}

type UserDetails struct {
	FirstName  string `gorm:"size:11" json:"firstName"`                                       // 手机号
	FamilyName string `gorm:"size:11" json:"familyName"`                                      // 手机号
	Phone      string `gorm:"size:11" json:"phone"`                                           // 手机号
	RoleId     int    `gorm:"-" json:"roleId"`                                                // 角色编码
	DeptId     int    `gorm:"-" json:"deptId"`                                                //部门编码
	PostionId  int    `gorm:"-" json:"PostionId"`                                             //职位编码
	Avatar     string `gorm:"size:255" json:"avatar"`                                         //头像
	Stated     int    `gorm:"type:enum('published', 'pending', 'deleted');default:'pending'"` //性别
	Email      string `gorm:"size:128" json:"email"`                                          //邮箱
}
