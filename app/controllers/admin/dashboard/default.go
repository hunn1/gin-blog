package dashboard

import (
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	ginview.HTML(c, http.StatusOK, "main/main", nil)
}
