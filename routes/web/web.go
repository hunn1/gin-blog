package web

import (
	"Kronos/app/Http/Controllers/Home"
	"Kronos/helpers"
	"Kronos/library/template"
	"github.com/gin-gonic/gin"
	html "html/template"
)

func RegisterWebRouter(router *gin.Engine) {
	// HTML 模板

	router.Static("resources/public", "./resources/public")
	router.HTMLRender = template.LoadTemplates("resources/views/home")
	router.SetFuncMap(html.FuncMap{
		// 注册模板函数
		"formatAsDate": helpers.FormatAsDate,
	})
	web := router.Group("/")
	{
		web.GET("/", Home.IndexApi)
	}

}
