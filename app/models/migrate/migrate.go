package models

import (
	"Kronos/app/models"
	"Kronos/library/databases"
	"github.com/sirupsen/logrus"
)

var Models = []interface{}{
	&models.Article{},
	&models.Admin{},
	&models.Permissions{},
	&models.RolePermission{},
	&models.Roles{},
	&models.UserRole{},
}

func AutoMigrate() {
	db := databases.GetDB()
	// 自动创建数据库
	if migerr := db.Set("gorm:table_options", "ENGINE=Innodb").AutoMigrate(Models...).Error; nil != migerr {
		logrus.Fatal("auto migrate tables failed: " + migerr.Error())
	}
}
