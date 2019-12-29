package admin

import (
	"Kronos/app/controllers/admin"
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

	group := router.Group("/admin", givMid)
	{
		//a := mysqladapter.NewAdapter("mysql", "root:@tcp(127.0.0.1:3306)/")
		//e := casbin.NewEnforcer("examples/basic_model.conf", a)
		//group.Use(middle.NewAuthorizer(e))
		group.GET("/", admin.ShowLogin)
	}
}
