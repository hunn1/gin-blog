package models

import "Kronos/library/databases"

type Tags struct {
	BaseModel
	Name string `gorm:"type:varchar(100);unique_index;not null;index:name;"`
}

func (t Tags) Get(cd string, vals []interface{}) (Tags, error) {
	if err := databases.DB.Model(&t).Where(cd, vals).First(&t).Error; err != nil {
		return t, err
	}
	return t, nil
}

func (t Tags) Lists(cd string, val []interface{}, ofst, limit int) (tgs []Tags, err error) {
	if err := databases.DB.Model(&t).Where(cd, val).Offset(ofst).Limit(limit).Find(&tgs).Error; err != nil {
		return tgs, err
	}
	return tgs, nil
}

func (t Tags) Update(data map[string]interface{}) {

}

func (t Tags) Create() {

}
