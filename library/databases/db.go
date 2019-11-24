package databases

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func InitDB(DbType, user, pass, dbname, charset, loc string) {
	dabs := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True&loc=%s", user, pass, dbname, charset, loc)
	db, _ := gorm.Open(DbType, dabs)
	defer db.Close()
}
