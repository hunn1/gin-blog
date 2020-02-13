package helpers

import "github.com/gin-gonic/gin"

func NewApiReturn(code int, msg string, data interface{}) *gin.H {
	return &gin.H{
		"code":    code,
		"message": msg,
		"data":    data,
	}
}

func NewApiRedirect(code int, msg string, redirectUrl string) *gin.H {

	return &gin.H{
		"code":        code,
		"message":     msg,
		"redirectUrl": redirectUrl,
	}
}
