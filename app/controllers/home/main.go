package home

import (
	"Kronos/app/models"
	"Kronos/helpers"
	"Kronos/library/databases"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func IndexApi(c *gin.Context) {

	//title := c.Query("title")
	//key := c.Query("keyword")
	Page, _ := strconv.ParseInt(c.DefaultQuery("p", "1"), 10, 0)
	PageSize, _ := strconv.ParseInt(c.DefaultQuery("limit", "1"), 10, 0)
	// 文章模型
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

	//查询分页的数据
	artCount.Offset((Page - 1) * PageSize).Limit(PageSize).Find(&artList)

	c.HTML(http.StatusOK, "main/main", gin.H{
		"title":   "Go Go Go !",
		"artList": artList,
	})
}
