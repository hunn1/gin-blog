package web

import (
	"Kronos/app/Http/Controllers/Home"
	"github.com/gin-gonic/gin"
)

func RegisterWebRouter(router *gin.Engine) {
	// HTML 模板
	router.LoadHTMLGlob("resources/views/home/**/*.html")
	router.Static("resources/public", "./resources/public")

	web := router.Group("/")
	{
		web.GET("/", Home.IndexApi)
	}
}
