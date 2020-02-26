package models

import (
	"Kronos/library/databases"
	"Kronos/library/page"
	"github.com/casbin/casbin/v2"
	"gopkg.in/go-playground/validator.v9"
)

// 管理员
type Admin struct {
	BaseModel
	Username string `gorm:"type:char(50);unique_index;not null;" binding:"required" validate:"min=6,max=32"`
	// 设置管理员账号 唯一并且不为空
	Password    string          `gorm:"size:255;not null;"`       // 设置字段大小为255
	LastLoginIp uint32          `gorm:"type:int(1);not null;"`    // 上次登录IP
	IsSuper     int             `gorm:"type:tinyint(1);not null"` // 是否超级管理员
	Enforcer    casbin.Enforcer `inject:""`
	Roles       []Roles         `json:"roles" gorm:"many2many:user_role;not null;"`
}

// Validate the fields.
func (u *Admin) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

func (u *Admin) GetByCount(whereSql string, vals []interface{}) (count int) {
	databases.DB.Model(u).Where(whereSql, vals).Count(&count)
	return
}

func (u *Admin) Lists(fields string, whereSql string, vals []interface{}, page *page.Pagination) []Admin {
	list := make([]Admin, page.Perineum)
	databases.DB.Model(u).Select(fields).Where(whereSql, vals).Offset(page.GetPage()).Limit(page.Perineum).Find(&list)
	return list
}

func (u Admin) Get(whereSql string, vals []interface{}) Admin {
	databases.DB.Model(&u).Where(whereSql, vals).First(&u)
	return u
}
