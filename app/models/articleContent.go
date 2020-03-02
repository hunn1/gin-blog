package models

type ArticleContent struct {
	BaseModel
	ArticleID uint64
	Body      string
}
