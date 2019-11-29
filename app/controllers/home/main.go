package home

import (
	"Kronos/app/models"
	"Kronos/library/databases"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexApi(c *gin.Context) {

	//title := c.Query("title")
	//key := c.Query("keyword")
	user := models.User{}
	first := databases.GetDB().First(&user)
	c.JSON(http.StatusOK, gin.H{
		"title": "Go Go Go !",
		"body":  first,
	})

}
