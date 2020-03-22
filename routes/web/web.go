package web

import (
	"Kronos/app/controllers/home"
	"Kronos/helpers"
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
		Funcs:        helpers.Builtins,
		DisableCache: true,
		Delims:       goview.Delims{},
	})
	web := router.Group("/", givMid)
	{
		web.GET("/", home.IndexApi)
		web.GET("/posts/:id", home.Posts)
		web.GET("/arh", home.Timeline)
		web.GET("/tags", home.TagLists)
		web.GET("/cate", home.CateLists)
	}

}
