package models

import (
	"Kronos/library/databases"
)

// 菜单权限
type Permissions struct {
	ID          uint64 `gorm:"primary_key" json:"id" structs:"id"`
	Title       string `gorm:"type:varchar(50);unique_index"` // 权限标题
	Description string `gorm:"type:char(64);"`                // 注解
	Slug        string `gorm:"type:varchar(50);"`             // 权限名称
	HttpPath    string `gorm:"type:text"`                     // URI路径
	Method      string `gorm:"type:char(10);"`
}

func (m *Permissions) GetMenus() []*Permissions {
	var allMenu []*Permissions
	databases.DB.Model(&allMenu).Find(&allMenu)
	return allMenu
}

func (m *Permissions) GetByCount(where string, values []interface{}) (count int) {
	databases.DB.Model(&m).Where(where, values).Count(&count)
	return
}

func (m *Permissions) Lists(fields string, where string, values []interface{}, offset, limit int) ([]Permissions, error) {
	list := make([]Permissions, limit)
	if err := databases.DB.Model(&list).Select(fields).Where(where, values).Offset(offset).Limit(limit).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (m Permissions) Get(where string, values []interface{}) (Permissions, error) {

	first := databases.DB.Model(&m).Where(where, values).First(&m)
	if first.Error != nil {
		return m, first.Error
	}
	return m, nil
}

func (m Permissions) Update(id int, data map[string]interface{}) error {

	find := databases.DB.Model(&m).Where("id = ?", id).Find(&m)
	if find.Error != nil {
		return find.Error
	}
	save := databases.DB.Model(&m).Update(data)

	if save.Error != nil {
		return save.Error
	}
	return nil
}

func (m Permissions) Create() error {

	create := databases.DB.Model(&m).Create(&m)
	if create.Error != nil {
		return create.Error
	}
	return nil
}

func EditPerRoles(id int) []int {
	var permission Permissions
	var role []Roles

	databases.DB.Model(&permission).Where("id = ? ", id, 0)
	pf := databases.GetPrefix()
	joins := " left join " + pf + "role_menu b on " + pf + "roles.id=b.role_id left join " + pf + "permissions c on c.id=b.permissions_id"
	databases.DB.Joins(joins).Where("c.id = ?", id).Find(&role)

	var roleList []int
	for _, v := range role {
		roleList = append(roleList, int(v.ID))
	}
	return roleList
}
