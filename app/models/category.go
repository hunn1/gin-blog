package models

import "Kronos/library/databases"

type Category struct {
	BaseModel
	Name string `gorm:"type:varchar(100);unique_index;not null;index:name;"`
}

func (c Category) Lists(fields, ws string, val []interface{}, ofst, limit int) (cate []Category, err error) {
	if err = databases.DB.Model(&c).Select(fields).Where(ws, val).Offset(ofst).Limit(limit).Find(&cate).Error; err != nil {
		return cate, err
	}
	return cate, nil
}

func (c Category) GetByCount(ws string, val []interface{}) (count int) {
	if err := databases.DB.Model(&c).Where(ws, val).Count(&count).Error; err != nil {
		return 0
	}
	return
}

func (c Category) Get(ws string, val []interface{}) (Category, error) {
	if err := databases.DB.Model(&c).Where(ws, val).Find(&c).Error; err != nil {
		return c, err
	}
	return c, nil
}

func (c Category) Update(id uint64, data map[string]interface{}) error {
	find := databases.DB.Model(&c).Where("id = ?", id).Find(&c)
	if find.Error != nil {
		return find.Error
	}
	save := find.Update(data)

	if save.Error != nil {
		return save.Error
	}
	return nil
}

func (c Category) Create() error {
	create := databases.DB.Model(&c).Create(&c)
	if create.Error != nil {
		return create.Error
	}
	return nil
}

func (c Category) Delete(id uint64) error {
	if err := databases.DB.Where("id = ?", id).Find(&c).Error; err != nil {
		return err
	}

	db := databases.DB.Model(&c).Where("id = ?", id).Delete(&c)
	if db.Error != nil {
		return db.Error
	}

	return nil
}

func (c Category) GetAll() (all []Category, err error) {
	if err := databases.DB.Model(&all).Find(&all).Error; err != nil {
		return nil, err
	}
	return
}
