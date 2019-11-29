package migrate

import (
	"Kronos/app/models"
)

var Models = []interface{}{
	&models.Article{}, &models.User{},
}

func AutoMigrate() {

	//if migerr := databases.GetDB().Set("gorm:table_options", "ENGINE=Innodb").AutoMigrate(Models...).Error; nil != migerr {
	//	logrus.Fatal("auto migrate tables failed: " + migerr.Error())
	//}
}
