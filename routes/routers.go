package routes

import (
	"Kronos/config"
	"Kronos/library/logs"
	"Kronos/routes/admin"
	"Kronos/routes/api"
	"Kronos/routes/web"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitRouters() *gin.Engine {
	_ = config.Init("./config/config.yaml")
	engine := gin.New()
	// 自定义中间件

	// 错误中间件
	//engine.Use(except.HandleErrors())

	// 使用日志
	engine.Use(gin.Logger())

	// 模式
	gin.SetMode(viper.GetString("runmode"))

	// 自定义日志存储
	engine.Use(logs.LoggerToFile())
	// 未找到路由 & 未找到操作方法
	//engine.NoRoute(func(context *gin.Context) {
	//	context.JSON(http.StatusNotFound, except.HandleErrors())
	//})
	//// 没找到操作方法
	//engine.NoMethod(func(context *gin.Context) {
	//	context.JSON(http.StatusNotFound, except.HandleErrors())
	//})
	engine.Static("resources/public", "./resources/public")
	// API 路由注册
	api.RegisterApiRouter(engine)
	// web 路由
	web.RegisterWebRouter(engine)
	admin.RegAdminRouter(engine)
	// ... 其他路由注册

	return engine
}
