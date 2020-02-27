package admins

import (
	"Kronos/app/controllers/admin"
	"Kronos/app/models"
	"Kronos/library/page"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"html/template"
)

type AdminsHandler struct {
	admin.AdminBaseHandler
}

func (a AdminsHandler) Lists(c *gin.Context) {

	all := a.AllParams(c)

	// 条件封装

	if all["username"] != nil {
		a.Where["username like"] = all["username"].(string) + "%"
	}

	build, vals, _ := models.WhereBuild(a.Where)
	// 列表页
	//
	var model = &models.Admin{}
	//查询总条数
	count := model.GetByCount(build, vals)
	// 分页
	pagination := page.NewPagination(c.Request, count, 10)
	// 查询数据绑定到列表slice
	fields := "username, last_login_ip, is_super,created_at"
	lists := model.Lists(fields, build, vals, pagination)

	ginview.HTML(c, 200, "admins/lists", gin.H{
		"pagination": template.HTML(pagination.Pages()),
		"total":      pagination.Total,
		"lists":      lists,
		"req":        all,
	})
}

// 添加或编辑
func (a AdminsHandler) ShowEdit(c *gin.Context) {
	query := a.AllParams(c)

	var model = models.Admin{}
	// 编辑
	if query["id"] != nil {
		a.Where["id"] = query["id"]
		build, vals, _ := models.WhereBuild(a.Where)
		model.Get(build, vals)
	}

	ginview.HTML(c, 200, "admins/edit", gin.H{
		"data": model,
		"req":  query,
	})
}

// 应用操作
func (a AdminsHandler) Apply(c *gin.Context) {

}

// 删除数据
func (a AdminsHandler) Delete(c *gin.Context) {

}
