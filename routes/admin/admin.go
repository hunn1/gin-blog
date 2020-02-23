package admin

import (
	"Kronos/app/controllers/admin/admins"
	"Kronos/app/controllers/admin/auth"
	"Kronos/app/controllers/admin/dashboard"
	"Kronos/app/controllers/admin/role"
	"Kronos/app/middle"
	"Kronos/helpers"
	"Kronos/library/casbin_adapter"
	"Kronos/library/casbin_helper"
	"Kronos/library/session"
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
)

func RegAdminRouter(router *gin.Engine) {

	// HTML 模板
	givMid := ginview.NewMiddleware(goview.Config{
		Root:         "resources/views/admin",
		Extension:    ".html",
		Master:       "layouts/master",
		Partials:     nil,
		Funcs:        helpers.TFMP,
		DisableCache: true,
		Delims:       goview.Delims{},
	})
	// Casbin
	e, err := casbin_adapter.InitAdapter()
	if err != nil {
		panic("无法初始化权限")
	}
	// 单独加载后台登录页
	router.LoadHTMLFiles("resources/views/admin/login/admin_login.html")

	// 分组使用母版内容
	router.Use(session.NewSessionStore())
	router.GET("/admin/login", auth.ShowLogin)
	router.POST("/admin/login", auth.Login)
	ntc := router.Group("/admin", givMid)
	{
		// 使用中间件认证
		ntc.Use(middle.AuthAdmin(e, casbin_helper.NotCheck("/admin/login", "/admin/logout")))
		// 初始化Session
		ntc.Use(session.AuthSessionMiddle())
		// 登出
		ntc.GET("logout", auth.Logout)

		// 后台面板
		ntc.GET("/", dashboard.Index)
		// 用户
		users := ntc.Group("admins")
		{
			users.GET("lists", admins.Lists)
			//users.GET("edit")
			//users.POST("apply")
			//users.POST("delete")
		}
		// 角色
		roles := ntc.Group("role")
		{
			roles.GET("lists", role.Lists)
			//roles.GET("edit")
			//roles.POST("apply")
			//roles.POST("delete")
		}

	}

}
