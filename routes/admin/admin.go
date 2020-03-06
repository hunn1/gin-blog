package admin

import (
	"Kronos/app/controllers/admin/admins"
	"Kronos/app/controllers/admin/articles"
	"Kronos/app/controllers/admin/auth"
	"Kronos/app/controllers/admin/category"
	"Kronos/app/controllers/admin/dashboard"
	"Kronos/app/controllers/admin/permissioins"
	"Kronos/app/controllers/admin/role"
	"Kronos/app/controllers/admin/tags"
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
			roles.GET("delete", roleHandler.Delete)
		}

		// 权限
		permission := ntc.Group("permission")
		{
			var permissionHandler = permissioins.PermissionHandler{}
			permission.GET("lists", permissionHandler.Lists)
			permission.GET("edit", permissionHandler.ShowEdit)
			permission.POST("apply", permissionHandler.Apply)
			permission.GET("delete", permissionHandler.Delete)
		}

		// 文章
		article := ntc.Group("article")
		{
			var artilceHandler = articles.ArticleHandler{}
			article.GET("lists", artilceHandler.Lists)
			article.GET("edit", artilceHandler.ShowEdit)
			article.POST("apply", artilceHandler.Apply)
			article.GET("delete", artilceHandler.Delete)
			article.GET("trash", artilceHandler.Trash)
			article.GET("force_delete", artilceHandler.ForceDelete)
		}

		// 文章
		tag := ntc.Group("tags")
		{
			var tagHandler = tags.TagHandler{}
			tag.GET("lists", tagHandler.Lists)
			tag.GET("edit", tagHandler.ShowEdit)
			tag.POST("apply", tagHandler.Apply)
			tag.GET("delete", tagHandler.Delete)
		}

		// 文章
		cate := ntc.Group("category")
		{
			var cateHandler = category.CateHandler{}
			cate.GET("lists", cateHandler.Lists)
			cate.GET("edit", cateHandler.ShowEdit)
			cate.POST("apply", cateHandler.Apply)
			cate.GET("delete", cateHandler.Delete)
		}

	}

}
