package models

//Role of system
type Role struct {
	Uid       string     `json:"uid,omitempty" gorm:"-"`
	Type      string     `gorm:"size:8" json:"dgraph.type,omitempty"`
	ID        int        `json:"id,omitempty" gorm:"primary_key;AUTO_INCREMENT"` // 角色编码
	Key       string     `json:"Key,omitempty" gorm:"size:128;"`                 //角色代码
	Name      string     `json:"Name,omitempty" gorm:"size:128;"`                // 角色名称
	Resources []Resource `json:"Resources,omitempty" gorm:"size:128;"`           // 角色拥有的资源
	ModelExtension
}
