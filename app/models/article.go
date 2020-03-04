package models

import (
	"Kronos/library/databases"
	"errors"
)

type Article struct {
	BaseModel

	Title          string           `gorm:"type:varchar(100);"`
	Keyword        string           `gorm:"type:varchar(100);"`
	Description    string           `gorm:"type:varchar(100);"`
	Thumb          string           `gorm:"size:255"` // 设置字段大小为255
	ArticleContent []ArticleContent `gorm:"foreignkey:article_id;association_foreignkey:id"`
	//Category    []Category       `gorm:"foreignkey:category_id"`
	//Tags        []Tags           `gorm:"foreignkey:tag_id"`
}

func (a Article) Count(where string, vals []interface{}) (int, error) {
	var count = 0
	err := databases.DB.Model(&a).Where(where, vals).Count(&count).Error
	if err != nil {
		return count, errors.New("暂无数据可查")
	}
	return count, nil
}
func (a Article) Lists(where string, vals []interface{}, offset, limit int) ([]Article, error) {
	list := make([]Article, limit)
	// .Preload("ArticleContent")
	err := databases.DB.Model(a).Where(where, vals).Offset(offset).Limit(limit).Find(&list)
	if err.Error != nil {
		return nil, errors.New("暂无数据可查")
	}
	return list, nil
}

func (a Article) Get(where string, vals []interface{}) (Article, error) {
	first := databases.DB.Model(a).Preload("ArticleContent").Where(where, vals).First(&a)
	if first.Error != nil {
		return a, first.Error
	}
	return a, nil
}

func (a Article) Update(id uint64, data []ArticleContent) error {

	first := databases.DB.Model(&a).Where("id = ?", id).First(&a)
	if first.Error != nil {
		return first.Error
	}

	association := databases.DB.Model(&a).Association("ArticleContent").Replace(data)
	if association.Error != nil {
		return association.Error
	}
	if err := databases.DB.Model(&a).Update(a).Error; err != nil {
		return err
	}
	return nil
}
func (a Article) Create(data []ArticleContent) error {
	err := databases.DB.Model(&a).Create(&a).Association("ArticleContent").Append(data).Error
	if err != nil {
		return err
	}
	return nil
}
