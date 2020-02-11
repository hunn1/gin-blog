package admin

import (
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Dashboard(c *gin.Context) {

	ginview.HTML(c, http.StatusOK, "main/main", nil)
}
