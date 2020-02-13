package helpers

import "github.com/gin-gonic/gin"

// Map 格式返回数据，统一调度返回
func NewApiReturn(code int, msg string, data interface{}) *gin.H {
	return &gin.H{
		"code":    code,
		"message": msg,
		"data":    data,
	}
}

// Map 格式返回状态 消息 跳转连接
func NewApiRedirect(code int, msg string, redirectUrl string) *gin.H {
	return &gin.H{
		"code":        code,
		"message":     msg,
		"redirectUrl": redirectUrl,
	}
}
