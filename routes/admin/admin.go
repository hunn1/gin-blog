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

	_, err = e.Enforce("alice", "data1", "read")
	// Modify the policy.
	// e.AddPolicy(...)
	// e.RemovePolicy(...)

	// Save the policy back to DB.
	err = e.SavePolicy()

	fmt.Println(e)
	if err != nil {
		panic("无法初始化权限")
	}

	ntc := router.Group("/admin", givMid)
	{
		ntc.Use(middle.AuthAdmin(e, casbin_helper.NotCheck("/admin/login")))
		ntc.GET("/", admin.ShowLogin)
		ntc.GET("login", admin.ShowLogin)
		ntc.GET("test", admin.TestC)
	}

}
