package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// 初始化Session 方式 始终会返回一个数据存储方式，默认为 cookie
func NewSessionStore() gin.HandlerFunc {
	typeOf := viper.GetString("session.type")
	secret := viper.GetString("session.secret")
	name := viper.GetString("session.name")
	// 判断
	switch typeOf {
	case "redis":
		store := redisStore(secret)
		return sessions.Sessions(name, store)

	case "cookie":
		fallthrough
	default:
		store := cookieStore(secret)
		return sessions.Sessions(name, store)
	}
}

// Redis
func redisStore(secret string) redis.Store {
	host := viper.GetString("redis.host")
	port := viper.GetString("redis.port")
	pass := viper.GetString("redis.pass")
	store, _ := redis.NewStore(10, "tcp", host+":"+port, pass, []byte(secret))
	return store
}

// Cookie
func cookieStore(secret string) cookie.Store {
	store := cookie.NewStore([]byte(secret))
	return store
}
