package articles

import (
	"Kronos/app/controllers/admin"
	"Kronos/app/models"
	"Kronos/library/page"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"html/template"
)

type ArticleHandler struct {
	admin.AdminBaseHandler
	model *models.Article
}

func (a ArticleHandler) Lists(c *gin.Context) {
	req := a.AllParams(c)
	getMap := a.GetMap(10)
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

func (a ArticleHandler) ShowEdit(c *gin.Context) {

}

func (a ArticleHandler) Apply(c *gin.Context) {

}

func (a ArticleHandler) Delete(c *gin.Context) {

}
