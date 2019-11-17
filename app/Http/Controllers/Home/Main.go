package Home

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexApi(c *gin.Context) {
	c.HTML(http.StatusOK, "main/index.html", gin.H{
		"title": "Go Go Go !",
	})

}
