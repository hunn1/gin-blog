package home

import (
	"Kronos/app/models"
	"Kronos/helpers"
	"Kronos/library/databases"
	"Kronos/library/page"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func IndexApi(c *gin.Context) {

	var PageSize int64 = 10
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

	ginview.HTML(c, http.StatusOK, "main/main", gin.H{
		"title":   "",
		"artList": artList,
		"page":    template.HTML(pagination.Pages()),
	})
}

func Posts(c *gin.Context) {
	c.JSON(200, "Todo")
}

func Timeline(c *gin.Context) {
	ginview.HTML(c, http.StatusOK, "main/timeline", gin.H{})
}
