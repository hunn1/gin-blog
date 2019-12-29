package casbin_adapter

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	adapter "github.com/casbin/gorm-adapter/v2"
	"github.com/spf13/viper"
)

// 初始化权限 数据库适配器
func InitAdapter() (*casbin.SyncedEnforcer, error) {

	host := viper.GetString("db.host")
	user := viper.GetString("db.user")
	pass := viper.GetString("db.pass")
	dbname := viper.GetString("db.dbname")
	charset := viper.GetString("db.charset")
	loc := viper.GetString("db.loc")
	native := viper.GetString("db.native")

	dabs := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=%s&allowNativePasswords=%s",
		user, pass, host,
		dbname, charset, loc,
		native,
	)

	//a, err := adapter.NewAdapterByDB(databases.DB)
	a, err := adapter.NewAdapter("mysql", dabs, true)
	if err != nil {
		return nil, fmt.Errorf("can not Init: %v", err.Error())
	}
	e, err := casbin.NewSyncedEnforcer("./config/rbac_model.conf", a)
	return e, err
}
