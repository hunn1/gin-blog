package admins

import (
	"Kronos/app/controllers/admin"
	"Kronos/app/models"
	"Kronos/library/apgs"
	"Kronos/library/page"
	"Kronos/library/password"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"html/template"
	"strconv"
)

type AdminsHandler struct {
	admin.AdminBaseHandler
}

func (a AdminsHandler) Lists(c *gin.Context) {

	all := a.AllParams(c)

	// 条件封装
	where := a.GetWhere(10)
	if all["username"] != nil {
		where["username like"] = all["username"].(string) + "%"
	}

	build, vals, _ := models.WhereBuild(where)
	// 列表页
	//
	var model = &models.Admin{}
	//查询总条数
	count := model.GetByCount(build, vals)
	// 分页
	pagination := page.NewPagination(c.Request, count, 10)
	// 查询数据绑定到列表slice
	fields := "id, username, last_login_ip, is_super,created_at"
	lists, _ := model.Lists(fields, build, vals, pagination)

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
		where := a.GetWhere(10)

		where["id"] = query["id"]
		build, vals, _ := models.WhereBuild(where)
		model, _ = model.Get(build, vals)

	}
	ginview.HTML(c, 200, "admins/edit", gin.H{
		"admin": model,
		"req":   query,
	})
}

// 应用操作
func (a AdminsHandler) Apply(c *gin.Context) {
	id := c.PostForm("id")
	parseInt, _ := strconv.ParseInt(id, 10, 64)
	var model = models.Admin{}
	if parseInt > 0 {
		passowrd := c.PostForm("passowrd")
		IsSuper := c.PostForm("IsSuper")
		where := a.GetWhere(1)
		where["id"] = parseInt
		build, vals, _ := models.WhereBuild(where)
		get, _ := model.Get(build, vals)
		get.Password, _ = password.Encrypt(passowrd)
		i, _ := strconv.ParseInt(IsSuper, 10, 0)
		get.IsSuper = int(i)
		update, err := get.Update()

		if err != nil {
			c.JSON(200, apgs.NewApiReturn(4003, "无法更新该数据", err))
		}
		if update == false {
			c.JSON(200, apgs.NewApiReturn(4004, "更新失败", nil))
		}
		c.JSON(200, apgs.NewApiReturn(200, "更新成功", update))
	} else {
		err := c.ShouldBind(&model)

		if err == nil {
			model.Password, _ = password.Encrypt(model.Password)
			create, err := model.Create()
			if err != nil {
				c.JSON(200, apgs.NewApiReturn(4003, "无法创建该数据", nil))
			}
			create.Password = ""
			c.JSON(200, apgs.NewApiReturn(200, "创建成功", create))
		}
	}
}

// 删除数据
func (a AdminsHandler) Delete(c *gin.Context) {
	id := c.Query("id")
	parseInt, _ := strconv.ParseInt(id, 10, 64)
	var mod = models.Admin{}
	if parseInt > 0 {
		where := a.GetWhere(10)
		where["id"] = parseInt
		build, vals, _ := models.WhereBuild(where)
		get, err := mod.Get(build, vals)
		if err != nil {
			c.JSON(200, apgs.NewApiReturn(4001, "该数据不存在或无法访问", nil))
		}
		b, err := get.Delete()
		if err != nil {
			c.JSON(200, apgs.NewApiReturn(4004, "无法删除该数据", nil))
		}
		c.JSON(200, apgs.NewApiReturn(200, "删除成功", b))

	}

}
