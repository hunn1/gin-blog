package home

import (
	"Kronos/app/models"
	"Kronos/helpers"
	"Kronos/library/databases"
	"Kronos/library/page"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"strconv"
)

func IndexApi(c *gin.Context) {

	var PageSize int64 = 10
	// 文章模型
	var artList []models.Article

	var count int32
	artCount := databases.DB.Model(models.Article{}).Count(&count)

	if artCount.Error != nil {
		helpers.Abort(c, "没有数据")
	}
	pagination := page.NewPagination(c.Request, int(count), int(PageSize))

	//查询分页的数据
	artCount.Offset((pagination.Page - 1) * PageSize).Limit(PageSize).Find(&artList)

	ginview.HTML(c, http.StatusOK, "main/main", gin.H{
		"title":   "",
		"artList": artList,
		"page":    template.HTML(pagination.Pages()),
	})
}

// 显示文章详情
func Posts(c *gin.Context) {
	var model models.Article
	var where = make(map[string]interface{}, 1)
	req_id := c.Param("id")
	id, _ := strconv.ParseInt(req_id, 10, 64)
	if id <= 0 {
		ginview.HTML(c, http.StatusNotFound, "err/404", gin.H{
			"message": "无法获取到ID",
		})
		return
	}

	where["id"] = id
	build, vals, err := models.WhereBuild(where)
	if err != nil {
		ginview.HTML(c, http.StatusNotFound, "err/msg", nil)
		return
	}
	get, err := model.Get(build, vals)
	if err != nil {
		ginview.HTML(c, http.StatusNotFound, "err/msg", gin.H{
			"msg": err,
		})
		return
	}
	ginview.HTML(c, http.StatusOK, "main/posts", gin.H{
		"data": get,
	})
}

// 显示时间戳归档
func Timeline(c *gin.Context) {
	ginview.HTML(c, http.StatusOK, "main/timeline", gin.H{})
}

// 标签列表
func TagLists(c *gin.Context) {
	var tags models.Tags
	allTags, _ := tags.GetAll()

	req_id := c.Param("id")
	id, _ := strconv.ParseInt(req_id, 10, 64)
	var art []models.Article
	if id > 0 {
		databases.DB.Model(&tags).First(&tags, id)
		databases.DB.Model(&tags).Preload("Tags").Preload("Category").Related(&art, "Article")
		//databases.DB.Model(&tags).Association("Article").Find(&art)
	}

	c.JSON(200, gin.H{
		"allTags": allTags,
		"art":     art,
	})
}

// 分类列表
func CateLists(c *gin.Context) {

	//req_id := c.Param("id")
	//id, _ := strconv.ParseInt(req_id, 10, 64)

	var cate models.Category
	allCate, _ := cate.GetAll()

	ginview.HTML(c, http.StatusOK, "main/cate", gin.H{
		"cateLists": allCate,
	})

}
