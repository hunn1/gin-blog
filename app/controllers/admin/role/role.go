package role

import (
	"Kronos/library/page"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"html/template"
)

func Lists(c *gin.Context) {
	page := page.NewPagination(c.Request, 100, 1).Pages()
	ginview.HTML(c, 200, "role/lists", gin.H{"page": template.HTML(page)})
}
