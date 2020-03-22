package models

//文章_标签
type ArticleTags struct {
	ArticleId string `gorm:"index" json:"article_id;"`
	TagsId    string `json:"tags_id"`
}
