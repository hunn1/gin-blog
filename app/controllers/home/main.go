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
	article := models.Article{}

	first := databases.DB.First(&article)
	c.JSON(http.StatusOK, gin.H{
		"title": "Go Go Go !",
		"body":  first,
	})
}
