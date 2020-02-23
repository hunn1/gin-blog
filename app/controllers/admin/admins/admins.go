package admins

import (
	"Kronos/app/models"
	"Kronos/library/databases"
	"Kronos/library/page"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"html/template"
)

func Lists(c *gin.Context) {

	list := make([]models.Admin, 10)

	find := databases.DB.Model(&list)
	var count int
	find.Count(&count)
	page := page.NewPagination(c.Request, count, 1)
	find.Select("username, last_login_ip, is_super,created_at").Offset(page.GetPage()).Limit(page.Perineum).Find(&list)
	//c.JSON(200, list)
	ginview.HTML(c, 200, "admins/lists", gin.H{
		"page":  template.HTML(page.Pages()),
		"total": page.Total,
		"lists": list,
		"req":   "",
	})
}
