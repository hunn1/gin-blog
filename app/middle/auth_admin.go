package middle

import (
	"Kronos/helpers"
	"Kronos/library/casbin_helper"
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
		userId := c.GetString("user_id")
		p := c.Request.URL.Path
		m := c.Request.Method
		if b, err := enforcer.Enforce(userId, p, m); err != nil {
			c.JSON(401, helpers.ApiReturn{Code: 401, Message: err.Error()})
			return
		} else if !b {
			c.JSON(401, helpers.ApiReturn{Code: 401, Message: err.Error()})
			return
		}
		c.Next()
	}
}
