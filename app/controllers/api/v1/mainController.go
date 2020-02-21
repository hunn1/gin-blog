package v1

import (
	"Kronos/library/apgs"
	"github.com/gin-gonic/gin"
)

func ShowApi(c *gin.Context) {
	var msg struct {
		Name    string `json:"user"`
		Message string
		Number  int
	}
	msg.Name = "Lena"
	msg.Message = "hey"
	msg.Number = 123
	apgsReturn := apgs.NewApiReturn(200, "我不知道呀", "msg")
	c.JSON(200, apgsReturn)
}
