package models

type Tags struct {
	BaseModel
	Name string `gorm:"type:varchar(100);unique_index;not null;index:name;"`
}
