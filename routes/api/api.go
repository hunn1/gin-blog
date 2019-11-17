package api

import (
	"Kronos/app/controllers/api/v1"
	"github.com/gin-gonic/gin"
)

func RegisterApiRouter(router *gin.Engine) {
	api := router.Group("api")
	api.GET("/test/index", v1.ShowApi)
	api.GET("/index", v1.ShowApi)
	api.GET("/cookie/set/:user_id", v1.ShowApi)

}
