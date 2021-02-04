package session

import (
	"encoding/json"
	"github.com/gin-contrib/sessions"

	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

const (
	UserKey = "userID"
)

// 初始化Session 方式 始终会返回一个数据存储方式，默认为 cookie
func NewSessionStore() gin.HandlerFunc {
	typeOf := viper.GetString("session.type")
	secret := viper.GetString("session.secret")
	name := viper.GetString("session.name")
	// 判断
	var store sessions.Store
	switch typeOf {
	case "redis":
		store = redisStore(secret)
		return sessions.Sessions(name, store)
	case "cookie":
		store = cookieStore(secret)
		return sessions.Sessions(name, store)
	default:
		store = cookieStore(secret)
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

// 登录Session 中间件
func AuthSessionMiddle() gin.HandlerFunc {

	return func(c *gin.Context) {
		session := sessions.Default(c)
		userId := session.Get(UserKey)
		if userId == nil {
			//ginview.HTML(
			//	c,
			//	http.StatusUnauthorized,
			//	"err/401",
			//	helpers.NewApiRedirect(200, "请登录后重试...", "/admin/login"),
			//)
			c.Redirect(302, "/admin/login")
			c.Abort()
			return
		}
		c.Set(UserKey, userId)
		c.Next()
	}
}

// 获取Session
func GetUserSession(c *gin.Context) uint {
	session := sessions.Default(c)
	userId := session.Get(UserKey)
	if userId == nil {
		return 0
	}
	return userId.(uint)
}

// 判断是否有Session
func HadSession(c *gin.Context) bool {
	session := sessions.Default(c)
	if userId := session.Get(UserKey); userId == nil {
		return false
	}
	return true
}

// 清除session
func ClearAuthSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	err := session.Save()
	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)
		c.Abort()
		return
	}
}

func GetSession(c *gin.Context, key interface{}) interface{} {
	session := sessions.Default(c)
	get := session.Get(key)
	return get
}

// 保存Session
func SaveSession(c *gin.Context, key interface{}, val interface{}) bool {
	session := sessions.Default(c)

	inrec, _ := json.Marshal(val)
	session.Set(key, inrec)
	err := session.Save()
	return err == nil
}
