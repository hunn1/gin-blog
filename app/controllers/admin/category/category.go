package category

import (
	"Kronos/app/controllers/admin"
	"Kronos/app/models"
	"Kronos/library/apgs"
	"Kronos/library/page"
	"encoding/json"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"html/template"
	"strconv"
)

type CateHandler struct {
	admin.AdminBaseHandler
}

func (p *CateHandler) Lists(c *gin.Context) {

	all := p.AllParams(c)

	// 条件封装
	where := p.GetMap(10)
	if all["name"] != nil {
		where["name like"] = all["name"].(string) + "%"
	}

	build, vals, _ := models.WhereBuild(where)
	// 列表页
	//
	var model = &models.Category{}
	//查询总条数
	count := model.GetByCount(build, vals)
	// 分页
	pagination := page.NewPagination(c.Request, count, 10)
	// 查询数据绑定到列表slice
	fields := "name,id,created_at, updated_at"
	lists, _ := model.Lists(fields, build, vals, pagination.GetPage(), pagination.Perineum)

	ginview.HTML(c, 200, "category/lists", gin.H{
		"pagination": template.HTML(pagination.Pages()),
		"total":      pagination.Total,
		"lists":      lists,
		"req":        all,
	})
}

func (p *CateHandler) ShowEdit(c *gin.Context) {
	query := p.AllParams(c)

	var model = models.Category{}
	// 编辑
	if query["id"] != nil {
		where := p.GetMap(10)

		where["id"] = query["id"]
		build, vals, _ := models.WhereBuild(where)
		model, _ = model.Get(build, vals)
	}

	ginview.HTML(c, 200, "category/edit", gin.H{
		"data": model,
		"req":  query,
	})
}

func (p *CateHandler) Apply(c *gin.Context) {

	var model = models.Category{}
	err := c.ShouldBind(&model)

	if err != nil {
		c.JSON(200, apgs.NewApiReturn(300, "无法获取到参数", nil))
		return
	}

	if model.ID > 0 {
		v := p.GetMap(10)
		marshal, _ := json.Marshal(model)
		_ = json.Unmarshal(marshal, &v)
		err = model.Update(model.ID, v)
		if err != nil {
			c.JSON(200, apgs.NewApiReturn(4003, "无法更新该数据", err))
			return
		}
		c.JSON(200, apgs.NewApiRedirect(200, "更新成功", "/admin/category/lists"))
		return

	} else {

		err := model.Create()
		if err != nil {
			c.JSON(200, apgs.NewApiReturn(4003, "无法创建该数据", nil))
			return
		}
		c.JSON(200, apgs.NewApiRedirect(200, "创建成功", "/admin/category/lists"))
		return
	}
}

func (p *CateHandler) Delete(c *gin.Context) {
	id := c.Query("id")
	parseInt, _ := strconv.ParseInt(id, 10, 64)
	var mod = models.Category{}
	if parseInt <= 0 {
		c.JSON(200, apgs.NewApiReturn(4004, "ID不能为0", nil))
		return
	}
	err := mod.Delete(uint64(parseInt))
	if err != nil {
		c.JSON(200, apgs.NewApiReturn(4004, "无法删除该数据", nil))
		return
	}
	c.JSON(200, apgs.NewApiRedirect(200, "删除成功", "/admin/admins/lists"))
}
