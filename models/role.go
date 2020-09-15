package models

//Role of system
type Role struct {
	ID        int        `json:"id" gorm:"primary_key;AUTO_INCREMENT"` // 角色编码
	Key       string     `json:"Key" gorm:"size:128;"`                 //角色代码
	Name      string     `json:"Name" gorm:"size:128;"`                // 角色名称
	Resources []Resource `json:"Resources" gorm:"size:128;"`           // 角色拥有的资源
	ModelExtension
}
