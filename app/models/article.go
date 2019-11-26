package models

type Article struct {
	BaseModel

	Title       string `gorm:"type:varchar(100);"`
	Keyword     string `gorm:"type:varchar(100);"`
	Description string `gorm:"type:varchar(100);"`
	Thumb       string `gorm:"size:255"` // 设置字段大小为255
}
