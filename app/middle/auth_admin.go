package middle

import (
	"Kronos/helpers"
	"Kronos/library/casbin_helper"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

// AuthAdmin 中间件
func AuthAdmin(enforcer *casbin.SyncedEnforcer, nocheck ...casbin_helper.DontCheckFunc) gin.HandlerFunc {

	return func(c *gin.Context) {
		if casbin_helper.DontCheck(c, nocheck...) {
			c.Next()
			return
		}
		userId := "1"

		p := c.Request.URL.Path
		m := c.Request.Method

		fmt.Println("UserID:" + userId)
		fmt.Println("Path:" + p)
		fmt.Println("Method:" + m)

		if b, err := enforcer.Enforce(userId, p, m); err != nil {
			c.JSON(401, helpers.NewApiReturn(401, err.Error(), nil))
			c.Abort()
			return
		} else if !b {
			c.JSON(401, helpers.NewApiReturn(401, "权限验证失败", b))
			c.Abort()
			return
		}
		c.Next()
	}
}
