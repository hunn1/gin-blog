package admin

import (
	"Kronos/app/controllers/admin/admins"
	"Kronos/app/controllers/admin/auth"
	"Kronos/app/controllers/admin/dashboard"
	"Kronos/app/controllers/admin/permissioins"
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
		Funcs:        helpers.Builtins,
		DisableCache: true,
		Delims:       goview.Delims{},
	})
	// Casbin
	Egor, err := casbin_adapter.InitAdapter()
	if err != nil {
		panic(err)
	}
	// 单独加载后台登录页
	router.LoadHTMLFiles("resources/views/admin/login/admin_login.html")

	// 分组使用母版内容
	router.Use(session.NewSessionStore())
	var auth = &auth.LoginHandler{}
	router.GET("/admin/login", auth.ShowLogin)
	router.POST("/admin/login", auth.Login)
	ntc := router.Group("/admin", givMid)
	{
		// 使用中间件认证
		ntc.Use(middle.AuthAdmin(Egor, casbin_helper.NotCheck("/admin/login", "/admin/logout")))
		// 初始化Session
		ntc.Use(session.AuthSessionMiddle())
		// 登出
		ntc.GET("logout", auth.Logout)

		// 后台面板
		ntc.GET("/", dashboard.Index)
		// 用户
		users := ntc.Group("admins")
		{
			var admHandler = admins.AdminsHandler{}
			users.GET("lists", admHandler.Lists)
			users.GET("edit", admHandler.ShowEdit)
			users.POST("apply", admHandler.Apply)
			users.GET("delete", admHandler.Delete)
		}
		// 角色
		roles := ntc.Group("role")
		{
			var roleHandler = role.RolesHandler{}
			roles.GET("lists", roleHandler.Lists)
			roles.GET("edit", roleHandler.ShowEdit)
			roles.POST("apply", roleHandler.Apply)
			roles.POST("delete", roleHandler.Delete)
		}

		// 角色
		permission := ntc.Group("permission")
		{
			var permissionHandler = permissioins.PermissionHandler{}
			permission.GET("lists", permissionHandler.Lists)
			permission.GET("edit", permissionHandler.ShowEdit)
			permission.POST("apply", permissionHandler.Apply)
			permission.POST("delete", permissionHandler.Delete)
		}

	}

}
