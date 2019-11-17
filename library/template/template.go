package template

import (
	"github.com/gin-contrib/multitemplate"
	"path/filepath"
	"strings"
)

func LoadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	// 读取layout文件
	articleLayouts, err := filepath.Glob(templatesDir + "/layouts/layout.html")
	if err != nil {
		panic(err.Error())
	}
	// 读取模板文件
	articles, err := filepath.Glob(templatesDir + "/*/*.html")
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our articleLayouts/ and articles/ directories
	for _, article := range articles {
		layoutCopy := make([]string, len(articleLayouts))
		copy(layoutCopy, articleLayouts)
		files := append(layoutCopy, article)
		// 转换为 模块@文件的形式，避免重复文件的问题
		filePath := filepath.Dir(article)
		r.AddFromFiles(strings.Split(filePath, "/")[3]+"@"+filepath.Base(article), files...)
	}
	return r
}
