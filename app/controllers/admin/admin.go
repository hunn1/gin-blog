package admin

import (
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"strings"
)

type AdminBaseHandler struct {
	Where map[string]interface{}
}

func (h *AdminBaseHandler) AllParams(c *gin.Context) map[string]interface{} {
	query := c.Request.URL.Query()
	queryMap := make(map[string]interface{}, 3)
	for key, value := range query {
		if len(value) > 0 {
			isFilter := strings.ContainsAny(key, "filter_")
			if isFilter != false {
				index := strings.Index(key, "_")
				if value[0] != "" {
					queryMap[key[index+1:len(key)]] = value[0]
				}
			}
		}

	}
	return queryMap

}

func (h *AdminBaseHandler) ShowError(c *gin.Context, url string) {
	ginview.HTML(c, 200, "err/redirect", gin.H{
		"wait": 3,
		"url":  url,
	})
	c.Abort()
}
