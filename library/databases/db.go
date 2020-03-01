package databases

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

var DB *gorm.DB

// 初始化DB
func InitDB() {
	// 数据库初始化
	DbType := viper.GetString("db.type")
	host := viper.GetString("db.host")
	user := viper.GetString("db.user")
	pass := viper.GetString("db.pass")
	dbname := viper.GetString("db.dbname")
	charset := viper.GetString("db.charset")
	loc := viper.GetString("db.loc")
	native := viper.GetString("db.native")
	prefix := viper.GetString("db.prefix")
	debug := viper.GetBool("db.debug")
	var err error
	dabs := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=%s&allowNativePasswords=%s",
		user, pass, host,
		dbname, charset, loc,
		native,
	)
	// 设置数据库连接数

	db, err := gorm.Open(DbType, dabs)
	if err != nil {
		logrus.Fatal("Cannot Connect : " + err.Error())
	}
	db.LogMode(debug)
	db.DB().SetMaxIdleConns(100)
	db.DB().SetMaxOpenConns(1000)
	db.DB().SetConnMaxLifetime(5 * time.Minute)
	//err := db.DB().Ping()

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return prefix + defaultTableName
	}
	DB = db

}

func GetPrefix() string {
	return viper.GetString("db.prefix")
}
func GetDB() *gorm.DB {
	return DB
}

// 好像不需要关闭数据库连接 先写着
func CloseDB() {
	if err := DB.Close(); nil != err {
		logrus.Fatal("Disconnect from database failed: " + err.Error())
	}
}
