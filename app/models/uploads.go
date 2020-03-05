package models

import (
	"Kronos/library/databases"
	"github.com/jinzhu/gorm"
)

// 管理员
type Uploads struct {
	BaseModel
	Resources string `gorm:"type:varchar(255); unique_index; not null; "`
}

func (u Uploads) GetByCount(whereSql string, vals []interface{}) (count int) {
	databases.DB.Model(u).Where(whereSql, vals).Count(&count)
	return
}

func (u Uploads) Lists(fields string, whereSql string, vals []interface{}, offset, limit int) ([]Uploads, error) {
	list := make([]Uploads, limit)
	find := databases.DB.Model(&u).Select(fields).Where(whereSql, vals).Offset(offset).Limit(limit).Find(&list)
	if find.Error != nil && find.Error != gorm.ErrRecordNotFound {
		return nil, find.Error
	}
	return list, nil
}

func (u Uploads) Get(whereSql string, vals []interface{}) (Uploads, error) {
	first := databases.DB.Model(&u).Where(whereSql, vals).First(&u)
	if first.Error != nil {
		return u, first.Error
	}
	return u, nil
}

func (u Uploads) GetById(id int) (Uploads, error) {
	first := databases.DB.Model(&u).Where("id = ?", id).First(&u)
	if first.Error != nil {
		return u, first.Error
	}
	return u, nil
}

func (u Uploads) Create(data map[string]interface{}) (*Uploads, error) {
	create := databases.DB.Model(&u).Create(&u)
	if create.Error != nil {
		return nil, create.Error
	}
	return &u, nil
}

func (u Uploads) Update(id int, data map[string]interface{}) error {

	find := databases.DB.Model(&u).Where("id = ?", id).Find(&u)
	if find.Error != nil {
		return find.Error
	}

	save := databases.DB.Model(&u).Update(data)

	if save.Error != nil {
		return save.Error
	}
	return nil
}

func (u Uploads) Delete(id int) (bool, error) {
	databases.DB.Where("id = ?", id).Find(&u)
	databases.DB.Model(&u).Association("Roles").Delete()
	db := databases.DB.Model(&u).Where("id = ?", id).Delete(&u)
	if db.Error != nil {
		return false, db.Error
	}

	return true, nil
}
