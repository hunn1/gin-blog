package V1

import (
	"github.com/gin-gonic/gin"
	"Kronos/helpers"
)

func ShowApi(c *gin.Context)  {
	var msg struct {
		Name    string `json:"user"`
		Message string
		Number  int
	}
	msg.Name = "Lena"
	msg.Message = "hey"
	msg.Number = 123
	apiReturn := helpers.ApiReturn{200, "我不知道呀", msg}
	c.JSON(200, apiReturn)
}

