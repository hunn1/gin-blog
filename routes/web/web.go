package web

import (
	"Kronos/app/controllers/home"
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
)

func RegisterWebRouter(router *gin.Engine) {
	// HTML 模板

	givMid := ginview.NewMiddleware(goview.Config{
		Root:         "resources/views/home",
		Extension:    ".html",
		Master:       "layouts/master",
		Partials:     nil,
		Funcs:        nil,
		DisableCache: true,
		Delims:       goview.Delims{},
	})
	web := router.Group("/", givMid)
	{
		web.GET("/", home.IndexApi)
	}

}
