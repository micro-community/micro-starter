package models

//Role of system
type Role struct {
	ID       int    `json:"id" gorm:"primary_key;AUTO_INCREMENT"` // 角色编码
	Key      string `json:"Key" gorm:"size:128;"`                 //角色代码
	Name     string `json:"Name" gorm:"size:128;"`                // 角色名称
	Admin    bool   `json:"admin" gorm:"size:4;"`                 // 管理员
	UpdateBy string `json:"updateBy" gorm:"size:128;"`            //
	AddedBy  string `json:"addedBy" gorm:"size:128;"`             // 资源最后的添加
}
