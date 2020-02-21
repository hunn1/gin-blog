package models

import (
	"Kronos/library/databases"
	"github.com/jinzhu/gorm"
)

// 角色
type Roles struct {
	ID          uint64        `gorm:"primary_key" json:"id" structs:"id"`
	Title       string        `gorm:"type:varchar(50);unique_index" json:"title"` // 角色标题
	Description string        `gorm:"type:char(64);" json:"description"`          // 角色注解
	Permissions []Permissions `json:"permissions" gorm:"many2many:role_menu;"`
}

// 按照ID查找
func FindByID(id int) (bool, error) {
	var role Roles
	err := databases.DB.Select("id").Where("id = ? ", id).First(&role).Error
	if err != nil {
		return false, err
	}
	if role.ID > 0 {
		return true, nil
	}
	return false, nil
}

// 依据传入的条件查找条数
func GetCount(maps interface{}) (int, error) {
	var count int
	if err := databases.DB.Model(&Roles{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// 获取角色列表
func GetRolesPage(pageNum int, pageSize int, maps interface{}) ([]*Roles, error) {
	var role []*Roles
	err := databases.DB.Preload("Permissions").Where(maps).Offset(pageNum).Limit(pageSize).Find(&role).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return role, nil
}

// 按照ID  获取角色
func GetRoleByID(id int) (*Roles, error) {
	var role Roles
	err := databases.DB.Preload("Permissions").Where("id = ?", id).First(&role).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &role, nil
}

// 确认角色名称是否已存在
func CheckRoleName(name string) (bool, error) {
	var role Roles
	err := databases.DB.Where("title=?", name).First(&role).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, nil
	}
	if role.ID > 0 {
		return true, nil
	}
	return false, nil
}

// 编辑角色
func EditRole(id int, data map[string]interface{}) error {
	var role []Roles
	var permsiss Permissions
	databases.DB.Where("id in (?)", data["permissions_id"].(int)).Find(&permsiss)

	err := databases.DB.Where("id = ?", id).Find(&role).Error
	if err != nil {
		return err
	}
	databases.DB.Model(&role).Association("Permissions").Replace(permsiss)
	databases.DB.Model(&role).Update(data)
	return nil
}

// 添加角色
func AddRole(data map[string]interface{}) (id int, err error) {
	role := Roles{
		Title:       data["title"].(string),
		Description: data["description"].(string),
	}
	var per Permissions
	databases.DB.Where("id in (?)", data["permissions_id"].(int)).Find(&per)
	err = databases.DB.Create(&role).Association("Permissions").Append(&per).Error
	if err != nil {
		return 0, err
	}
	return int(role.ID), nil
}

// 删除角色
func DeleteRole(id int) error {
	var role Roles
	databases.DB.Where("id = ?", id).First(&role)
	databases.DB.Model(&role).Association("Permissions").Delete()
	err := databases.DB.Where("id = ?", id).Delete(&role).Error
	if err != nil {
		return err
	}
	return nil
}

// 删除所有角色
func CleanRole() error {
	//Unscoped 方法可以物理删除记录
	if err := databases.DB.Unscoped().Where("deleted_on != ? ", 0).Delete(&Roles{}).Error; err != nil {
		return err
	}

	return nil
}

// 获取所有角色
func GetRolesAll() ([]*Roles, error) {
	var role []*Roles
	err := databases.DB.Preload("Permissions").Find(&role).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return role, nil
}
