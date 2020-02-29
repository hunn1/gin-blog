package role

import (
	"Kronos/app/controllers/admin"
	"Kronos/app/models"
	"Kronos/library/apgs"
	"Kronos/library/page"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"html/template"
	"strconv"
)

type RolesHandler struct {
	admin.AdminBaseHandler
}

func (r *RolesHandler) Lists(c *gin.Context) {
	params := r.AllParams(c)
	where := r.GetWhere(10)

	if params["title"] != nil {
		where["title like"] = params["title"].(string) + "%"
	}

	build, vals, _ := models.WhereBuild(where)
	var model = models.Roles{}
	count, _ := model.GetCount(build, vals)
	page := page.NewPagination(c.Request, count, 10)
	list, _ := model.GetRolesPage(build, vals, page)

	ginview.HTML(c, 200, "role/lists", gin.H{
		"pagination": template.HTML(page.Pages()),
		"lists":      list,
		"total":      count,
		"req":        params,
	})
}

func (r *RolesHandler) ShowEdit(c *gin.Context) {
	params := r.AllParams(c)
	var model = models.Roles{}
	if params["id"] != nil {
		where := r.GetWhere(10)
		where["id"] = params["id"]
		build, vals, _ := models.WhereBuild(where)
		model, _ = model.Get(build, vals)
	}
	ginview.HTML(c, 200, "role/edit", gin.H{
		"role":        model,
		"req":         params,
		"permissions": "",
	})
}

func (r *RolesHandler) Apply(c *gin.Context) {
	id := c.PostForm("id")

	parseInt, _ := strconv.ParseInt(id, 10, 64)
	var model = models.Roles{}
	if parseInt > 0 {
		err := c.ShouldBind(&model)
		v := r.GetWhere(10)

		err = model.EditRole(int(parseInt), v)

		if err != nil {
			c.JSON(200, apgs.NewApiReturn(4003, "无法更新该数据", err))
			return
		}

		c.JSON(200, apgs.NewApiRedirect(200, "更新成功", "/admin/role/lists"))
		return

	} else {
		err := c.ShouldBind(&model)

		if err == nil {

			c.JSON(200, apgs.NewApiRedirect(200, "创建成功", "/admin/role/lists"))
			return
		}
	}
}

func (r *RolesHandler) Delete(c *gin.Context) {

}
