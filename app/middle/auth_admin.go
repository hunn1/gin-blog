package middle

import (
	"Kronos/helpers"
	"Kronos/library/casbin_helper"
	"github.com/casbin/casbin/v2"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// AuthAdmin 中间件
func AuthAdmin(enforcer *casbin.SyncedEnforcer, nocheck ...casbin_helper.DontCheckFunc) gin.HandlerFunc {

	return func(c *gin.Context) {
		if casbin_helper.DontCheck(c, nocheck...) {
			c.Next()
			return
		}
		userId := "1"

		p := strings.ToLower(c.Request.URL.Path)
		m := strings.ToLower(c.Request.Method)

		//fmt.Println("UserID  v0 :" + userId)
		//fmt.Println("Path v1 :" + p)
		//fmt.Println("Method v2 :" + m)

		var b bool
		var err error
		if b, err = enforcer.Enforce(userId, p, m); err != nil {
			// TODO 判断是是否为调试模式
			// TODO 调试模式下 判断 异步，同步 返回 JSON HTML
			//c.JSON(403, helpers.NewApiReturn(401, err.Error(), b))
			//c.AbortWithStatus(403)
			ginview.HTML(c, http.StatusForbidden, "err/403", gin.H{"errMsg": err.Error()})
			c.Abort()
			return
		}
		if !b {
			//c.JSON(401, helpers.NewApiReturn(401, "权限验证失败", b))
			//c.Abort()
			//fmt.Println("Check:" + strconv.FormatBool(b))
			//c.Redirect(302, "/admin/login")
			ginview.HTML(c, http.StatusUnauthorized, "err/401", helpers.NewApiRedirect(200, "无权限访问该内容", "/admin/login"))
			c.Abort()
			return
		}
		c.Next()
	}
}
