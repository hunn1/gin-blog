package home

import (
	"Kronos/app/models"
	"Kronos/helpers"
	"Kronos/library/databases"
	"Kronos/library/page"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"strconv"
)

func IndexApi(c *gin.Context) {

	PageSize, _ := strconv.ParseInt(c.DefaultQuery("limit", "1"), 10, 0)
	// 文章模型
	var artList []models.Article

	var count int32
	artCount := databases.DB.Model(models.Article{}).Count(&count)
	if artCount.Error != nil {
		helpers.Abort(c, "没有数据")
	}
	pagination := page.NewPagination(c.Request, int(count), int(PageSize))

	//查询分页的数据
	artCount.Offset((pagination.Page - 1) * PageSize).Limit(PageSize).Find(&artList)

	c.HTML(http.StatusOK, "main/main", gin.H{
		"title":   "Go Go Go !" + strconv.Itoa(int(pagination.Page-1)),
		"artList": artList,
		"page":    template.HTML(pagination.Pages()),
	})
}
