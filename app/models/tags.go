package models

import "Kronos/library/databases"

type Tags struct {
	BaseModel
	Name    string    `gorm:"type:varchar(100);unique_index;not null;index:name;"`
	Article []Article `gorm:"many2many:article_tags;" json:",omitempty"`
}

func (t Tags) GetByCount(cd string, val []interface{}) (count int) {
	if err := databases.DB.Model(&t).Where(cd, val).Count(&count).Error; err != nil {
		return 0
	}
	return
}

func (t Tags) Get(cd string, vals []interface{}) (Tags, error) {
	if err := databases.DB.Model(&t).Where(cd, vals).First(&t).Error; err != nil {
		return t, err
	}
	return t, nil
}

func (t Tags) Lists(fields, cd string, val []interface{}, ofst, limit int) (tgs []Tags, err error) {
	if err := databases.DB.Model(&t).Select(fields).Where(cd, val).Offset(ofst).Limit(limit).Find(&tgs).Error; err != nil {
		return tgs, err
	}
	return tgs, nil
}

func (t Tags) Update(id uint64, data map[string]interface{}) error {
	first := databases.DB.Model(&t).Where("id = ?", id).First(&t)
	if first.Error != nil {
		return first.Error
	}
	update := first.Update(data)
	if update.Error != nil {
		return update.Error
	}
	return nil
}

func (t Tags) Create() error {
	create := databases.DB.Model(&t).Create(&t)
	if create.Error != nil {
		return create.Error
	}
	return nil
}

func (t Tags) Delete(id uint64) error {
	first := databases.DB.Model(&t).Where("id = ?", id).First(&t)
	if first.Error != nil {
		return first.Error
	}
	if err := first.Delete(&t).Error; err != nil {
		return err
	}
	return nil
}

func (t Tags) GetAll() (all []Tags, err error) {
	if err := databases.DB.Model(&all).Find(&all).Error; err != nil {
		return nil, err
	}
	return
}
