package articles

import (
	"Kronos/app/controllers/admin"
	"Kronos/app/models"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
)

type ArticleHandler struct {
	admin.AdminBaseHandler
	model *models.Article
}

func (a ArticleHandler) Lists(c *gin.Context) {
	// a.model.Lists("", []interface{})
	ginview.HTML(c, 200, "admins/lists", gin.H{})
}
