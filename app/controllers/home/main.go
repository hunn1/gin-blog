package home

import (
	"Kronos/app/models"
	"Kronos/helpers"
	"Kronos/library/databases"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexApi(c *gin.Context) {

	//title := c.Query("title")
	//key := c.Query("keyword")

	var artList []models.Article
	//err := databases.DB.First(&article)
	//if err.Error != nil {
	//	helpers.Abort(c, "Test Message")
	//}
	var count int32
	artCount := databases.DB.Model(models.Article{}).Count(&count)
	if artCount.Error != nil {
		helpers.Abort(c, "没有数据")
	}
	Page := 1
	PageSize := 5
	//查询分页的数据
	artCount.Offset((Page - 1) * PageSize).Limit(PageSize).Find(&artList)

	c.HTML(http.StatusOK, "main/main", gin.H{
		"title":   "Go Go Go !",
		"artList": artList,
	})
}
