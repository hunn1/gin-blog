package Exceptions

import (
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"net/http"
)

// 错误信息
func HandleErrors() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			// 致命错误 捕获恢复
			if err := recover(); err != nil {
				var (
					errMsg     string
					mysqlError *mysql.MySQLError
					ok         bool
				)
				// assert type
				if errMsg, ok = err.(string); ok {
					c.JSON(http.StatusInternalServerError, gin.H{
						"code":    500,
						"message": "内部错误." + errMsg,
					})
					return
				} else if mysqlError, ok = err.(*mysql.MySQLError); ok {
					c.JSON(http.StatusInternalServerError, gin.H{
						"code":    500,
						"message": "Sorry , We Lost Database " + mysqlError.Error(),
					})
					return
				} else {
					c.JSON(http.StatusInternalServerError, gin.H{
						"code":    500,
						"message": "Oops , Something Went Wrong Message...",
					})
					return
				}
			}
		}()
		c.Next()
	}
}
