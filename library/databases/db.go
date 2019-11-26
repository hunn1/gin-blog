package databases

import (
	"Kronos/app/models"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"time"
)

var db *gorm.DB

// 初始化DB
func InitDB(DbType, host, user, pass, dbname, charset, loc, native string) {
	var err error
	dabs := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=%s&allowNativePasswords=%s", user, pass, host, dbname, charset, loc, native)
	// 设置数据库连接数

	db, err = gorm.Open(DbType, dabs)
	if err != nil {
		logrus.Fatal("Cannot Connect : " + err.Error())
	}
	if migerr := db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(models.Models...).Error; nil != migerr {
		logrus.Fatal("auto migrate tables failed: " + err.Error())
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(5 * time.Minute)

}

// 好像不需要关闭数据库连接 先写着
func CloseDB() {
	if err := db.Close(); nil != err {
		logrus.Fatal("Disconnect from database failed: " + err.Error())
	}
}
