package databases

import (
	"Kronos/app/migrate"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"time"
)

var db *gorm.DB

// 初始化DB
func InitDB(DbType, host, user, pass, dbname, charset, loc, native, prefix string) {
	var err error
	dabs := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=%s&allowNativePasswords=%s",
		user, pass, host,
		dbname, charset, loc,
		native,
	)
	// 设置数据库连接数

	db, err = gorm.Open(DbType, dabs)
	if err != nil {
		logrus.Fatal("Cannot Connect : " + err.Error())
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(5 * time.Minute)
	// 自动创建数据库
	migrate.AutoMigrate()
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return prefix + defaultTableName
	}
}
func GetDB() *gorm.DB {
	return db
}

// 好像不需要关闭数据库连接 先写着
func CloseDB() {
	if err := db.Close(); nil != err {
		logrus.Fatal("Disconnect from database failed: " + err.Error())
	}
}
