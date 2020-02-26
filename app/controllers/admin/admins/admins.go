package admins

import (
	"Kronos/app/models"
	"Kronos/helpers"
	"Kronos/library/page"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"html/template"
)

func Lists(c *gin.Context) {

	all := helpers.GetMapFilterQuery(c.Request.URL.Query())

	// 条件封装
	var where = make(map[string]interface{})
	if all["username"] != nil {

		where["username like"] = all["username"].(string) + "%"
	}

	build, vals, _ := models.WhereBuild(where)
	// 列表页
	//
	var admin = &models.Admin{}
	//查询总条数
	count := admin.GetByCount(build, vals)
	// 分页
	pagination := page.NewPagination(c.Request, count, 10)
	// 查询数据绑定到列表slice
	fields := "username, last_login_ip, is_super,created_at"
	lists := admin.Lists(fields, build, vals, pagination)

	ginview.HTML(c, 200, "admins/lists", gin.H{
		"pagination": template.HTML(pagination.Pages()),
		"total":      pagination.Total,
		"lists":      lists,
		"req":        all,
	})
}

// 添加或编辑
func ShowEdit(c *gin.Context) {
	query := helpers.GetMapFilterQuery(c.Request.URL.Query())
	where := make(map[string]interface{})
	var admin = models.Admin{}
	// 编辑
	if query["id"] != nil {
		where["id"] = query["id"]
		build, vals, _ := models.WhereBuild(where)
		admin.Get(build, vals)
	}

	ginview.HTML(c, 200, "admins/edit", gin.H{
		"data": admin,
		"req":  query,
	})
}

// 应用操作
func Apply(c *gin.Context) {

}

// 删除数据
func Delete(c *gin.Context) {

}
