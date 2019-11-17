package web

import (
	"Kronos/app/Http/Controllers/Home"
	"Kronos/library/template"
	"github.com/gin-gonic/gin"
)

func RegisterWebRouter(router *gin.Engine) {
	// HTML 模板

	router.Static("resources/public", "./resources/public")
	router.HTMLRender = template.LoadTemplates("resources/views/home")
	web := router.Group("/")
	{
		web.GET("/", Home.IndexApi)
	}

}
