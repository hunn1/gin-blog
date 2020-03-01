package role

import (
	"Kronos/app/controllers/admin"
	"Kronos/app/models"
	"Kronos/library/apgs"
	"Kronos/library/page"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"html/template"
	"strings"
)

type RolesHandler struct {
	admin.AdminBaseHandler
}

func (r *RolesHandler) Lists(c *gin.Context) {
	params := r.AllParams(c)
	where := r.GetMap(10)

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
	mapd := make(map[string]uint64, 10)

	if params["id"] != nil {
		where := r.GetMap(10)
		where["id"] = params["id"]
		build, vals, _ := models.WhereBuild(where)
		model, _ = model.Get(build, vals)

		for _, i2 := range model.Permissions {
			mapd[i2.HttpPath] = i2.ID
		}

	}

	var permission = models.Permissions{}
	menus := permission.GetMenus()
	v := make(map[string][]interface{}, 0)

	for _, menu := range menus {
		split := strings.Split(menu.HttpPath, "/")
		if _, ok := v[split[2]]; ok {
			v[split[2]] = append(v[split[2]], menu)
		} else {
			vds := make([]interface{}, 0)
			v[split[2]] = append(vds, menu)
		}

	}
	//for _, i2 := range v {
	//	fmt.Println(i2)
	//	if len(i2) > 0 {
	//		for _, i3 := range i2 {
	//			fmt.Println(i3.(*models.Permissions).Title)
	//		}
	//	}
	//}

	ginview.HTML(c, 200, "role/edit", gin.H{
		"role":        model,
		"req":         params,
		"permissions": v,
		"mapd":        mapd,
	})
}

func (r *RolesHandler) Apply(c *gin.Context) {

	var model = models.Roles{}
	err := c.ShouldBind(&model)
	if err != nil {
		c.JSON(200, apgs.NewApiReturn(4003, "无法获取数据", err))
		return
	}
	permissionIds := c.PostFormArray("permission[]")

	if model.ID > 0 {
		v := r.GetMap(10)
		v["permissions_id"] = permissionIds

		err = model.EditRole(int(model.ID), v)
		if err != nil {
			c.JSON(200, apgs.NewApiReturn(4003, "无法更新数据", err.Error()))
			return
		}
		model.LoadPolicy(int(model.ID))
		c.JSON(200, apgs.NewApiRedirect(200, "更新成功", "/admin/role/lists"))
		return

	} else {

		_, err := model.AddRole(map[string]interface{}{
			"title":          model.Title,
			"description":    model.Description,
			"permissions_id": permissionIds,
		})
		if err != nil {
			c.JSON(200, apgs.NewApiReturn(4003, "无法创建数据", nil))
			return
		}
		model.LoadPolicy(int(model.ID))
		c.JSON(200, apgs.NewApiRedirect(200, "创建成功", "/admin/role/lists"))
		return
	}
}

func (r *RolesHandler) Delete(c *gin.Context) {

}
