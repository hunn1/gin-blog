package articles

import (
	"Kronos/app/controllers/admin"
	"Kronos/app/models"
	"Kronos/helpers"
	"Kronos/library/apgs"
	"Kronos/library/page"
	"encoding/json"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"html/template"
	"strconv"
)

type ArticleHandler struct {
	admin.AdminBaseHandler
	model models.Article
}

func (a ArticleHandler) Lists(c *gin.Context) {
	req := a.AllParams(c)
	getMap := a.GetMap(10)
	if req["title"] != nil {
		getMap["title like"] = req["title"].(string) + "%"
	}
	build, vals, _ := models.WhereBuild(getMap)
	total, _ := a.model.Count(build, vals)
	p := page.NewPagination(c.Request, total, 10)
	lists, _ := a.model.Lists(build, vals, p.GetPage(), p.Perineum)

	ginview.HTML(c, 200, "article/lists", gin.H{
		"lists": lists,
		"req":   req,
		"total": total,
		"page":  template.HTML(p.Pages()),
	})
}
func (a ArticleHandler) Trash(c *gin.Context) {
	req := a.AllParams(c)
	getMap := a.GetMap(10)
	if req["title"] != nil {
		getMap["title like"] = req["title"].(string) + "%"
	}
	getMap["deleted_at"] = models.IsNotNull
	build, vals, _ := models.WhereBuild(getMap)
	total, _ := a.model.Count(build, vals)
	p := page.NewPagination(c.Request, total, 10)
	lists, _ := a.model.Trash(build, vals, p.GetPage(), p.Perineum)

	ginview.HTML(c, 200, "article/trash", gin.H{
		"lists": lists,
		"req":   req,
		"total": total,
		"page":  template.HTML(p.Pages()),
	})
}

func (a ArticleHandler) ShowEdit(c *gin.Context) {
	params := a.AllParams(c)

	if params["id"] != nil {
		getMap := a.GetMap(3)
		getMap["id"], _ = strconv.ParseInt(params["id"].(string), 10, 0)
		build, vals, _ := models.WhereBuild(getMap)
		a.model, _ = a.model.Get(build, vals)

	}
	cateModel := models.Category{}
	allcate, _ := cateModel.GetAll()
	tagModel := models.Tags{}
	allTags, _ := tagModel.GetAll()
	curr_tags, _ := json.Marshal(a.model.Tags)
	curr_cate, _ := json.Marshal(a.model.Category)
	ginview.HTML(c, 200, "article/edit", gin.H{
		"data":      a.model,
		"req":       params,
		"cate":      allcate,
		"tags":      allTags,
		"curr_tags": string(curr_tags),
		"curr_cate": string(curr_cate),
	})
}

func (a ArticleHandler) Apply(c *gin.Context) {

	array := c.PostFormMap("content")
	categoryIds := c.PostFormArray("category[]")
	tagIds := c.PostFormArray("tags[]")
	form := c.PostFormArray("content[]")

	err := c.ShouldBind(&a.model)
	if err != nil {
		c.JSON(200, apgs.NewApiReturn(3003, "无法获取到数据", nil))
		return
	}

	file, err := helpers.UploadFile(c, "cover_photo")
	if err != nil {
		c.JSON(200, apgs.NewApiReturn(3001, err.Error(), nil))
		return
	}

	if file != "" {
		a.model.Thumb = file
	}

	var artc = make([]models.ArticleContent, 0)
	if a.model.ID > 0 {
		for id, i2 := range array {
			parseInt, _ := strconv.ParseInt(id, 10, 64)
			artc = append(artc, models.ArticleContent{ID: uint64(parseInt), ArticleID: 1, Body: i2})
		}
		for _, i2 := range form {
			if i2 != "" {
				artc = append(artc, models.ArticleContent{Body: i2, ArticleID: a.model.ID})
			}
		}
		a.model.ArticleContent = artc
		err := a.model.Update(a.model.ID, map[string]interface{}{
			"category_ids": categoryIds,
			"tag_ids":      tagIds,
		})
		if err != nil {
			c.JSON(200, apgs.NewApiReturn(4005, "无法更新该文章数据", err.Error()))
			return
		}
	} else {
		for _, i2 := range form {
			artc = append(artc, models.ArticleContent{Body: i2})
		}
		a.model.ArticleContent = artc
		err := a.model.Create(map[string]interface{}{
			"category_ids": categoryIds,
			"tag_ids":      tagIds,
		})
		if err != nil {
			c.JSON(200, apgs.NewApiReturn(4004, "无法创建该文章数据", err.Error()))
			return
		}
	}
	c.JSON(200, apgs.NewApiReturn(300, "操作成功", nil))
	return

}

func (a ArticleHandler) Delete(c *gin.Context) {
	id := c.Query("id")
	parseInt, _ := strconv.ParseInt(id, 10, 64)
	var mod = models.Article{}
	if parseInt <= 0 {
		c.JSON(200, apgs.NewApiReturn(4004, "ID不能为0", nil))
		return
	}
	err := mod.Delete(uint64(parseInt))
	if err != nil {
		c.JSON(200, apgs.NewApiReturn(4004, "无法删除该数据", nil))
		return
	}
	c.JSON(200, apgs.NewApiRedirect(200, "删除成功", "/admin/article/lists"))
	return
}

func (a ArticleHandler) ForceDelete(c *gin.Context) {
	id := c.Query("id")
	parseInt, _ := strconv.ParseInt(id, 10, 64)
	var mod = models.Article{}
	if parseInt <= 0 {
		c.JSON(200, apgs.NewApiReturn(4004, "ID不能为0", nil))
		return
	}
	err := mod.ForceDelete(uint64(parseInt))
	if err != nil {
		c.JSON(200, apgs.NewApiReturn(4004, "无法删除该数据", nil))
		return
	}
	c.JSON(200, apgs.NewApiRedirect(200, "删除成功", "/admin/article/trash"))
	return
}
