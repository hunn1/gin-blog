package template

import (
	"fmt"
	"github.com/gin-contrib/multitemplate"
	"path/filepath"
)

func LoadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	articleLayouts, err := filepath.Glob(templatesDir + "/layouts/layout.html")
	if err != nil {
		panic(err.Error())
	}

	articles, err := filepath.Glob(templatesDir + "/*/*.html")
	fmt.Println(articles)
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our articleLayouts/ and articles/ directories
	for _, article := range articles {
		layoutCopy := make([]string, len(articleLayouts))
		copy(layoutCopy, articleLayouts)
		files := append(layoutCopy, article)
		r.AddFromFiles(filepath.Base(article), files...)
	}
	return r
}
