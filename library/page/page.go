package page

import (
	"fmt"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

//Pagination 分页器
type Pagination struct {
	Request  *http.Request
	Total    int
	Perineum int
	Page     int64
}

//NewPagination 新建分页器
func NewPagination(req *http.Request, total int, perineum int) *Pagination {
	queryParams := req.URL.Query()
	//从当前请求中获取page
	page := queryParams.Get("page")
	if page == "" {
		page = "1"
	}
	p, _ := strconv.ParseInt(page, 10, 0)

	return &Pagination{
		Request:  req,
		Total:    total,
		Perineum: perineum,
		Page:     p,
	}
}
func (p *Pagination) GetPage() int {
	result := 0
	page := p.Page
	if page > 0 {
		result = int((page - 1) * int64(p.Perineum))
	}

	return result
}

// 获取当前页 开始记录
func (p *Pagination) FirstItem() int64 {
	return p.Page * int64(p.Perineum)
}

// 获取当前页结束记录
func (p *Pagination) LastItem() int64 {
	return p.FirstItem() + int64(p.Perineum)
}

//Pages 渲染生成html分页标签
func (p *Pagination) Pages() string {

	pageful := int(p.Page)
	if pageful == 0 {
		return ""
	}

	//计算总页数
	var totalPageNum = int(math.Ceil(float64(p.Total) / float64(p.Perineum)))

	//首页链接
	var firstLink string
	//上一页链接
	var prevLink string
	//下一页链接
	var nextLink string
	//末页链接
	var lastLink string
	//中间页码链接
	var pageLinks []string

	//首页和上一页链接
	if pageful > 1 {
		firstLink = fmt.Sprintf(`<li  class="paginate_button "><a href="%s">首页</a></li>`, p.pageURL("1"))
		prevLink = fmt.Sprintf(`<li class="paginate_button "><a href="%s">上一页</a></li>`, p.pageURL(strconv.Itoa(pageful-1)))
	} else {
		firstLink = `<li class="paginate_button disabled"><a href="#">首页</a></li>`
		prevLink = `<li class="paginate_button disabled"><a href="#">上一页</a></li>`
	}

	//末页和下一页
	if pageful < totalPageNum {
		lastLink = fmt.Sprintf(`<li class="paginate_button "><a href="%s">末页</a></li>`, p.pageURL(strconv.Itoa(totalPageNum)))
		nextLink = fmt.Sprintf(`<li class="paginate_button "><a href="%s">下一页</a></li>`, p.pageURL(strconv.Itoa(pageful+1)))
	} else {
		lastLink = `<li class="paginate_button disabled"><a href="#">末页</a></li>`
		nextLink = `<li class="paginate_button disabled"><a href="#">下一页</a></li>`
	}

	//生成中间页码链接
	pageLinks = make([]string, 0, 10)
	startPos := pageful - 3
	endPos := pageful + 3
	if startPos < 1 {
		endPos = endPos + int(math.Abs(float64(startPos))) + 1
		startPos = 1
	}
	if endPos > totalPageNum {
		endPos = totalPageNum
	}
	for i := startPos; i <= endPos; i++ {
		var s string
		if i == pageful {
			s = fmt.Sprintf(`<li class="paginate_button active"><a href="%s">%d</a></li>`, p.pageURL(strconv.Itoa(i)), i)
		} else {
			s = fmt.Sprintf(`<li class="paginate_button "><a href="%s">%d</a></li>`, p.pageURL(strconv.Itoa(i)), i)
		}
		pageLinks = append(pageLinks, s)
	}

	return fmt.Sprintf(`<ul class="">%s%s%s%s%s</ul>`, firstLink, prevLink, strings.Join(pageLinks, ""), nextLink, lastLink)
}

//pageURL 生成分页url
func (p *Pagination) pageURL(page string) string {
	//基于当前url新建一个url对象
	u, _ := url.Parse(p.Request.URL.String())
	q := u.Query()
	q.Set("page", page)
	u.RawQuery = q.Encode()
	return u.String()
}
