package models

type ResourceCatalog int

const (
	Device ResourceCatalog = iota
	Person
	Location
	Organization
	System
	UIMenu
	FunctionModule
)

type Resource struct {
	ID       int    `json:"id" gorm:"primary_key;AUTO_INCREMENT"` // 资源编码
	Key      string `json:"Key" gorm:"size:128;"`                 //资源代码
	Name     string `json:"Name" gorm:"size:128;"`                // 资源名称
	TenantID int    `json:"TenantId" gorm:"size:128;"`            // 租户ID ,是否是属于某个租户
	UpdateBy string `json:"updateBy" gorm:"size:128;"`            // 资源的更新时间
	AddedBy  string `json:"addedBy" gorm:"size:128;"`             // 资源最后的添加
	//Operations []Operation `json:"operations"`
}
