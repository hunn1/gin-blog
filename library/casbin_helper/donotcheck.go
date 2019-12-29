package casbin_helper

import (
	"github.com/gin-gonic/gin"
)

type DontCheckFunc func(ctx *gin.Context) bool

// NotCheck 指定路由不用检查
func NotCheck(prefixes ...string) DontCheckFunc {
	return func(c *gin.Context) bool {
		path := c.Request.URL.Path
		pathLen := len(path)

		for _, p := range prefixes {
			if pl := len(p); pathLen >= pl && path[:pl] == p {
				return true
			}
		}
		return false
	}
}

// Check 指定路由检查
func Check(prefixes ...string) DontCheckFunc {
	return func(c *gin.Context) bool {
		path := c.Request.URL.Path
		pathLen := len(path)

		for _, p := range prefixes {
			if pl := len(p); pathLen >= pl && path[:pl] == p {
				return false
			}
		}
		return true
	}
}

// DontCheck 不检查函数
func DontCheck(c *gin.Context, skippers ...DontCheckFunc) bool {
	for _, skipper := range skippers {
		if skipper(c) {
			return true
		}
	}
	return false
}
