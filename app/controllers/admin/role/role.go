package role

import (
	"Kronos/app/models"
	"Kronos/library/databases"
	"Kronos/library/page"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"html/template"
)

func Lists(c *gin.Context) {

	page := page.NewPagination(c.Request, 100, 10)

	list := make([]models.Roles, 10)
	databases.DB.Model(&list).Offset(0).Limit(10).Find(&list)

	ginview.HTML(c, 200, "role/lists", gin.H{
		"page": template.HTML(page.Pages()),
		"list": list,
	})
}
