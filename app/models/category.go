package models

type Category struct {
	BaseModel
	Name string `gorm:"type:varchar(100);unique_index;not null;index:name;"`
}
