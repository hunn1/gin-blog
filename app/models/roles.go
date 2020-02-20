package models

type Roles struct {
	ID          uint64 `gorm:"primary_key" json:"id" structs:"id"`
	Title       string `gorm:"type:varchar(50);unique_index" json:"title"` // 角色标题
	Description string `gorm:"type:char(64);" json:"description"`          // 角色注解
}
