package models

import (
	"Kronos/library/casbin_adapter"
	"Kronos/library/databases"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"gopkg.in/go-playground/validator.v9"
)

// 管理员
type Admin struct {
	BaseModel
	Username string `gorm:"type:char(50); unique_index;not null;"  validate:"min=6,max=32"`
	// 设置管理员账号 唯一并且不为空
	Password    string `gorm:"size:255;not null;"  `     // 设置字段大小为255
	LastLoginIp uint32 `gorm:"type:int(1);not null;"`    // 上次登录IP
	IsSuper     int    `gorm:"type:tinyint(1);not null"` // 是否超级管理员

	Roles []Roles `json:"roles" gorm:"many2many:admin_role;not null;"`
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

func (u Admin) Lists(fields string, whereSql string, vals []interface{}, offset, limit int) ([]Admin, error) {
	list := make([]Admin, limit)
	find := databases.DB.Preload("Roles").Model(&u).Select(fields).Where(whereSql, vals).Offset(offset).Limit(limit).Find(&list)
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

func (u Admin) GetById(id int) (Admin, error) {
	first := databases.DB.Preload("Roles").Model(&u).Where("id = ?", id).First(&u)
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

func (u Admin) Update(id int, data map[string]interface{}) error {
	var role = make([]Roles, 10)
	if err := databases.DB.Where("id in (?)", data["role_id"]).Find(&role).Error; err != nil {
		return errors.New("无法找到该角色")
	}

	find := databases.DB.Model(&u).Where("id = ?", id).Find(&u)
	if find.Error != nil {
		return find.Error
	}

	databases.DB.Model(&u).Association("Roles").Replace(role)
	save := databases.DB.Model(&u).Update(data)

	if save.Error != nil {
		return save.Error
	}
	return nil
}

func (u Admin) Delete(id int) (bool, error) {
	databases.DB.Where("id = ?", id).Find(&u)
	databases.DB.Model(&u).Association("Roles").Delete()
	db := databases.DB.Model(&u).Where("id = ?", id).Delete(&u)
	if db.Error != nil {
		return false, db.Error
	}
	casbin_adapter.GetEnforcer().DeleteUser(u.Username)
	return true, nil
}

// LoadPolicy 加载用户权限策略
func (u *Admin) LoadPolicy(id int) error {

	admin, err := u.GetById(id)
	if err != nil {
		return err
	}
	casbin_adapter.GetEnforcer().DeleteRolesForUser(admin.Username)

	for _, ro := range admin.Roles {
		casbin_adapter.GetEnforcer().AddRoleForUser(admin.Username, ro.Title)
	}
	fmt.Println("更新角色权限关系", casbin_adapter.GetEnforcer().GetGroupingPolicy())
	return nil
}

func (a Admin) GetUsersAll() ([]*Admin, error) {
	var admin []*Admin
	err := databases.DB.Model(&a).Preload("Roles").Find(&admin).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return admin, nil
}

// LoadAllPolicy 加载所有的用户策略
func (a *Admin) LoadAllPolicy() error {
	admins, err := a.GetUsersAll()
	if err != nil {
		return err
	}
	for _, admin := range admins {
		if len(admin.Roles) != 0 {
			err = a.LoadPolicy(int(admin.ID))
			if err != nil {
				return err
			}
		}
	}
	fmt.Println("角色权限关系", casbin_adapter.GetEnforcer().GetGroupingPolicy())
	return nil
}
