package models

type Permissions struct {
	ID          uint64 `gorm:"primary_key" json:"id" structs:"id"`
	Title       string `gorm:"type:varchar(50);unique_index"` // 权限标题
	Description string `gorm:"type:char(64);"`                // 注解
	Slug        string `gorm:"type:varchar(50);"`             // 权限名称
	HttpPath    string `gorm:"type:text"`                     // URI路径
}
