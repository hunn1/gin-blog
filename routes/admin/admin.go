package admin

import (
	"Kronos/app/controllers/admin"
	"Kronos/app/middle"
	"Kronos/library/casbin_adapter"
	"Kronos/library/casbin_helper"
	"fmt"
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
		Funcs:        nil,
		DisableCache: true,
		Delims:       goview.Delims{},
	})
	// Casbin
	e, err := casbin_adapter.InitAdapter()
	fmt.Println(e)
	if err != nil {
		panic("无法初始化权限")
	}
	router.Use(middle.AuthAdmin(e, casbin_helper.Check("/admin/")))

	ntc := router.Group("/admin", givMid)
	{
		ntc.GET("/", admin.ShowLogin)
		ntc.GET("login", admin.ShowLogin)
		ntc.GET("test", admin.TestC)
	}

}
