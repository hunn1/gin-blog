package models

import (
	"Kronos/library/databases"
	"Kronos/library/page"
	"errors"
	"github.com/casbin/casbin/v2"
	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v9"
)

// 管理员
type Admin struct {
	BaseModel
	Username string `gorm:"type:char(50); unique_index;not null;" form:"username" binding:"required" validate:"min=6,max=32"`
	// 设置管理员账号 唯一并且不为空
	Password    string          `gorm:"size:255;not null;" form:"password" binding:"required" ` // 设置字段大小为255
	LastLoginIp uint32          `gorm:"type:int(1);not null;"`                                  // 上次登录IP
	IsSuper     int             `gorm:"type:tinyint(1);not null"`                               // 是否超级管理员
	Enforcer    casbin.Enforcer `inject:""`
	Roles       []Roles         `json:"roles" gorm:"many2many:user_role;not null;"`
}

// Validate the fields.
func (u *Admin) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

func (u Admin) GetByCount(whereSql string, vals []interface{}) (count int) {
	databases.DB.Model(u).Where(whereSql, vals).Count(&count)
	return
}

func (u Admin) Lists(fields string, whereSql string, vals []interface{}, page *page.Pagination) ([]Admin, error) {
	list := make([]Admin, page.Perineum)
	find := databases.DB.Preload("Roles").Model(&u).Select(fields).Where(whereSql, vals).Offset(page.GetPage()).Limit(page.Perineum).Find(&list)
	if find.Error != nil && find.Error != gorm.ErrRecordNotFound {
		return nil, find.Error
	}
	return list, nil
}

func (u Admin) Get(whereSql string, vals []interface{}) (Admin, error) {
	first := databases.DB.Preload("Roles").Model(&u).Where(whereSql, vals).First(&u)
	if first.Error != nil {
		return u, first.Error
	}
	return u, nil
}

func (u Admin) Create(data map[string]interface{}) (*Admin, error) {
	var role = make([]Roles, 10)
	databases.DB.Where("id in (?)", data["role_id"]).Find(&role)
	create := databases.DB.Model(&u).Create(&u).Association("Roles").Append(role)
	if create.Error != nil {
		return nil, create.Error
	}
	return &u, nil
}

func (u Admin) Update(id int, data map[string]interface{}) (bool, error) {
	var role = make([]Roles, 10)
	if err := databases.DB.Where("id in (?)", data["role_id"]).Find(&role).Error; err != nil {
		return false, errors.New("无法找到该角色")
	}

	find := databases.DB.Model(&u).Where("id = ?", id).Find(&u)
	if find.Error != nil {
		return false, find.Error
	}

	databases.DB.Model(&u).Association("Roles").Replace(role)
	save := databases.DB.Model(&u).Update(data)

	if save.Error != nil {
		return false, save.Error
	}
	return true, nil
}

func (u Admin) Delete(id int) (bool, error) {
	databases.DB.Where("id = ?", id).Find(&u)
	databases.DB.Model(&u).Association("Roles").Delete()
	db := databases.DB.Model(&u).Where("id = ?", id).Delete(&u)
	if db.Error != nil {
		return false, db.Error
	}
	return true, nil
}
