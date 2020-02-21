package models

import (
	"github.com/casbin/casbin/v2"
	"gopkg.in/go-playground/validator.v9"
)

// 管理员
type Admin struct {
	BaseModel
	Username string `gorm:"type:char(50);unique_index;not null;" binding:"required" validate:"min=6,max=32"`
	// 设置管理员账号 唯一并且不为空
	Password    string          `gorm:"size:255;not null;"`       // 设置字段大小为255
	LastLoginIp int32           `gorm:"type:int(1);not null;"`    // 上次登录IP
	IsSuper     int             `gorm:"type:tinyint(1);not null"` // 是否超级管理员
	Enforcer    casbin.Enforcer `inject:""`
	Roles       []Roles         `json:"roles" gorm:"many2many:user_role;not null;"`
}

// Validate the fields.
func (u *Admin) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
