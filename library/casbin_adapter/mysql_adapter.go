package casbin_adapter

import (
	"Kronos/library/databases"
	"fmt"
	"github.com/casbin/casbin/v2"
	adapter "github.com/casbin/gorm-adapter/v2"
	"github.com/spf13/viper"
)

var Enfocer *casbin.SyncedEnforcer

// 初始化权限 数据库适配器
func InitAdapter() (*casbin.SyncedEnforcer, error) {

	if databases.DB == nil {
		databases.InitDB()
	}
	a, err := adapter.NewAdapterByDBUsePrefix(databases.DB, viper.GetString("db.prefix"))
	if err != nil {
		return nil, fmt.Errorf("can not Init: %v", err.Error())
	}
	e, err := casbin.NewSyncedEnforcer("./config/rbac_model.conf", a)
	if err != nil {
		return nil, fmt.Errorf("can not Init: %v", err.Error())
	}
	// 开启AutoSave机制
	e.EnableAutoSave(true)
	_ = e.BuildRoleLinks()
	enableLog := viper.GetBool("casbin.debug")
	e.EnableLog(enableLog)
	// 10秒重新加载一次权限
	//e.StartAutoLoadPolicy(10 * time.Second)
	//e.EnableAutoBuildRoleLinks(true)
	// 因为开启了AutoSave机制，现在内存中的改变会同步回写到持久层中
	//e.AddPolicy("admin", "test", "test")
	Enfocer = e
	return e, err
}

func GetEnforcer() *casbin.SyncedEnforcer {
	return Enfocer
}
