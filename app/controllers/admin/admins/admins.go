package admins

import (
	"Kronos/app/models"
	"Kronos/helpers"
	"Kronos/library/databases"
	"Kronos/library/page"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"html/template"
)

type AdminHandler struct {
}

func Lists(c *gin.Context) {

	all := helpers.GetMapFilterQuery(c.Request.URL.Query())

	// 条件封装
	var where = make(map[string]interface{})
	if all["username"] != nil {
		where["username like"] = all["username"]
	}

	build, vals, _ := models.WhereBuild(where)
	// 列表页
	list := make([]models.Admin, 10)
	find := databases.DB.Model(&list)
	// 总条数
	var count int
	find.Where(build, vals).Count(&count)
	// 分页
	page := page.NewPagination(c.Request, count, 1)
	// 查询数据绑定到列表slice
	find.Select("username, last_login_ip, is_super,created_at").Where(build, vals).Offset(page.GetPage()).Limit(page.Perineum).Find(&list)

	ginview.HTML(c, 200, "admins/lists", gin.H{
		"page":  template.HTML(page.Pages()),
		"total": page.Total,
		"lists": list,
		"req":   all,
	})
}
