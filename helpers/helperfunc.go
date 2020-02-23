package helpers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"time"
)

var TFMP = template.FuncMap{
	"showStatus": ShowStatus,
	"showtime":   ShowTime,
}

func ShowStatus(ts interface{}, on string, on2 string) (tsf string) {
	switch ts {
	case ts.(int):
		if ts == 0 {
			tsf = on
		} else if ts == 1 {
			tsf = on2
		}
	case ts.(string):
		if ts == "0" {
			tsf = on
		} else if ts == "1" {
			tsf = on2
		}
	}
	return tsf
}

func ShowTime(t time.Time, format string) string {
	return t.Format(format)
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d%02d/%02d", year, month, day)
}

func Abort(c *gin.Context, message string) {
	c.HTML(http.StatusInternalServerError, "/home/err/500.html", gin.H{
		"message": message,
	})
}
