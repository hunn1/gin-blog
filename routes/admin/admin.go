package admin

import (
	"Kronos/app/controllers/admin"
	"Kronos/app/middle"
	"Kronos/library/casbin_adapter"
	"Kronos/library/casbin_helper"
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
)

func RegAdminRouter(router *gin.Engine) {

	//	// HTML 模板
	givMid := ginview.NewMiddleware(goview.Config{
		Root:         "resources/views/admin",
		Extension:    ".html",
		Master:       "layouts/master",
		Partials:     nil,
		Funcs:        nil,
		DisableCache: true,
		Delims:       goview.Delims{},
	})

	e, err := casbin_adapter.InitAdapter()
	if err != nil {
		panic("无法初始化权限")
	}
	router.Use(middle.AuthAdmin(e, casbin_helper.NotCheck("/admin/login")))
	group := router.Group("/admin", givMid)
	{

		group.GET("/login", admin.ShowLogin)
	}
}
