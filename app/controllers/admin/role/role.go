package role

import (
	"Kronos/app/controllers/admin"
	"github.com/gin-gonic/gin"
)

type RolesHandler struct {
	admin.AdminBaseHandler
}

func (r *RolesHandler) Lists(c *gin.Context) {
	//params := r.AllParams(c)
	//where := r.GetWhere(10)
	//
	//if params["title"] != nil {
	//	where["title like"] = params["title"].(string) + "like"
	//}
	//
	//build, vals, _ := models.WhereBuild(where)
	//var model = models.Roles{}
	//count, _ := model.GetCount(build, vals)
	//
	//page := page.NewPagination(c.Request, count, 10)
	//list, _ := model.GetRolesAll(build, vals, page)
	//
	//ginview.HTML(c, 200, "role/lists", gin.H{
	//	"pagination": template.HTML(page.Pages()),
	//	"lists":      list,
	//	"total":      count,
	//	"req":        params,
	//})
}
