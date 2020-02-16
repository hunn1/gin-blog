package models

// 管理员
type Admin struct {
	BaseModel
	Username    string `gorm:"type:varchar(50);unique_index"` // 设置管理员账号 唯一并且不为空
	Password    string `gorm:"size:255"`                      // 设置字段大小为255
	LastLoginIp int32  `gorm:"type:int(1)"`                   // 上次登录IP
	IsSuper     int    `gorm:"type:tinyint(1)"`               // 是否超级管理员
}
