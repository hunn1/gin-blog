package api

import (
	"github.com/gin-gonic/gin"
	"Kronos/app/Http/Controllers/Api/V1"
)

func RegisterApiRouter(router *gin.Engine)  {
	api := router.Group("api")
	api.GET("/test/index", V1.ShowApi)
	api.GET("/index", V1.ShowApi)
	api.GET("/cookie/set/:user_id", V1.ShowApi)

}