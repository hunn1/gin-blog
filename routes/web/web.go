package web

import (
	"Kronos/app/controllers/home"
	"Kronos/helpers"
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	html "html/template"
)

func RegisterWebRouter(router *gin.Engine) {
	// HTML 模板

	router.Static("resources/public", "./resources/public")
	//router.HTMLRender = template.LoadTemplates("resources/views/home")
	router.HTMLRender = ginview.New(goview.Config{
		Root:         "resources/views/home",
		Extension:    ".html",
		Master:       "layouts/master",
		Partials:     nil,
		Funcs:        nil,
		DisableCache: true,
		Delims:       goview.Delims{},
	})
	router.SetFuncMap(html.FuncMap{
		// 注册模板函数
		"formatAsDate": helpers.FormatAsDate,
	})
	web := router.Group("/")
	{
		web.GET("/", home.IndexApi)
	}

}
