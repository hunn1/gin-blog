package models

import (
	"Kronos/library/databases"
	"errors"
)

type Article struct {
	BaseModel

	Title       string           `gorm:"type:varchar(100);"`
	Keyword     string           `gorm:"type:varchar(100);"`
	Description string           `gorm:"type:varchar(100);"`
	Thumb       string           `gorm:"size:255"` // 设置字段大小为255
	AllContents []ArticleContent `gorm:"foreignkey:article_id"`
}

func (a *Article) Count(where string, vals []interface{}) (int, error) {
	var count = 0
	err := databases.DB.Model(&a).Where(where, vals).Count(&count).Error
	if err != nil {
		return count, errors.New("暂无数据可查")
	}
	return count, nil
}
func (a *Article) Lists(where string, vals []interface{}, offset, limit int) ([]Article, error) {
	list := make([]Article, limit)
	err := databases.DB.Model(&a).Preload("AllContents").Where(where, vals).Offset(offset).Limit(limit).Find(&list)
	if err.Error != nil {
		return nil, errors.New("暂无数据可查")
	}
	return list, nil
}
