package admins

import (
	"Kronos/app/controllers/admin"
	"Kronos/app/models"
	"Kronos/library/apgs"
	"Kronos/library/page"
	"Kronos/library/password"
	"encoding/json"
	"fmt"
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
	var marshal []byte
	var model = models.Admin{}
	// 编辑
	if query["id"] != nil {
		where := a.GetWhere(10)

		where["id"] = query["id"]
		build, vals, _ := models.WhereBuild(where)
		model, _ = model.Get(build, vals)
		marshal, _ = json.Marshal(model.Roles)

		fmt.Println()
	}
	var role = models.Roles{}
	allRoles, _ := role.GetRolesAll()

	ginview.HTML(c, 200, "admins/edit", gin.H{
		"admin":   model,
		"req":     query,
		"roles":   allRoles,
		"marshal": string(marshal),
	})
}

// 应用操作
func (a AdminsHandler) Apply(c *gin.Context) {
	id := c.PostForm("id")
	roleId := c.PostFormArray("role_id[]")
	parseInt, _ := strconv.ParseInt(id, 10, 64)
	var model = models.Admin{}
	if parseInt > 0 {
		passowrd := c.PostForm("passowrd")
		IsSuper := c.PostForm("IsSuper")
		v := a.GetWhere(10)

		v["passowrd"], _ = password.Encrypt(passowrd)
		v["is_super"] = IsSuper
		v["role_id"] = roleId
		update, err := model.Update(int(parseInt), v)

		if err != nil {
			c.JSON(200, apgs.NewApiReturn(4003, "无法更新该数据", err))
			return
		}
		if update == false {
			c.JSON(200, apgs.NewApiReturn(4004, "更新失败", nil))
			return
		}

		c.JSON(200, apgs.NewApiRedirect(200, "更新成功", "/admin/admins/lists"))
		return

	} else {
		err := c.ShouldBind(&model)

		if err == nil {
			model.Password, _ = password.Encrypt(model.Password)
			v := a.GetWhere(10)
			v["role_id"] = roleId
			create, err := model.Create(v)
			if err != nil {
				c.JSON(200, apgs.NewApiReturn(4003, "无法创建该数据", nil))
				return
			}
			create.Password = ""
			c.JSON(200, apgs.NewApiRedirect(200, "创建成功", "/admin/admins/lists"))
			return
		}
	}
}

// 删除数据
func (a AdminsHandler) Delete(c *gin.Context) {
	id := c.Query("id")
	parseInt, _ := strconv.ParseInt(id, 10, 64)
	var mod = models.Admin{}
	if parseInt > 0 {
		_, err := mod.Delete(int(parseInt))
		if err != nil {
			c.JSON(200, apgs.NewApiReturn(4004, "无法删除该数据", nil))
			return
		}
		c.JSON(200, apgs.NewApiRedirect(200, "删除成功", "/admin/admins/lists"))
		return

	} else {
		c.JSON(200, apgs.NewApiReturn(4004, "ID不能为0", nil))
		return
	}

}
