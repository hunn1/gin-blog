package models

type ArticleContent struct {
	ID        uint64 `gorm:"primary_key" json:"id" structs:"id"`
	ArticleID uint64 `gorm:"not null;"`
	Body      string `gorm:"not null;"`
}
