package models

type UserRole struct {
	UserID uint64 `gorm:"primary_key;auto_increment:false;"`
	RoleID string `gorm:"primary_key;type:varchar(50);auto_increment:false;"` // 角色标题
}
