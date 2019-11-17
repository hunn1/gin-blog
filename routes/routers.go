package routes

import (
	"Kronos/Exceptions"
	"Kronos/config"
	"Kronos/library/logs"
	"Kronos/routes/api"
	"Kronos/routes/web"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouters() *gin.Engine {
	engine := gin.New()
	// 自定义中间件

	// 错误中间件
	engine.Use(Exceptions.HandleErrors())

	// 使用日志
	engine.Use(gin.Logger())
	_ = config.Init("./config/config.yaml")
	// 自定义日志存储
	engine.Use(logs.LoggerToFile())
	// 未找到路由 & 未找到操作方法
	engine.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, Exceptions.HandleErrors())
	})
	// 没找到操作方法
	engine.NoMethod(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, Exceptions.HandleErrors())
	})

	// API 路由注册
	api.RegisterApiRouter(engine)
	// web 路由
	web.RegisterWebRouter(engine)
	// ... 其他路由注册

	return engine
}
