package models

type RolePermission struct {
	RoleID       uint64 `gorm:"primary_key;auto_increment:false;"`
	PermissionID string `gorm:"primary_key;type:varchar(50);auto_increment:false;"` // 角色标题
}
