package models

import "Kronos/library/databases"

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

//func GetRolesPage(pageNum int, pageSize int, maps interface{}) ([]*Roles, error) {
//
//	return
//}

//func (r *Roles) Add() (id int, err error) {
//	role := map[string]interface{}{
//		"title":       r.Title,
//		"id":          r.ID,
//		"description": r.Description,
//	}
//
//}
