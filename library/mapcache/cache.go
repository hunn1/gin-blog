package mapcache

import (
	"fmt"
	"github.com/allegro/bigcache"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

var (
	DB          *gorm.DB
	GlobalCache *bigcache.BigCache
)

// 缓存
//https://neojos.com/blog/2018/08-19-%E6%9C%AC%E5%9C%B0%E7%BC%93%E5%AD%98bigcache/
func init() {
	// Initialize cache
	var err error
	GlobalCache, err = bigcache.NewBigCache(bigcache.DefaultConfig(30 * time.Minute))
	if err != nil {
		panic(fmt.Errorf("failed to initialize cahce: %w", err))
	}
}
