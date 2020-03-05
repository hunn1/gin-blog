package helpers

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"io"
	"net"
	"net/http"
	"os"
	"time"
)

var Builtins = template.FuncMap{
	"showStatus": ShowStatus,
	"showtime":   ShowTime,
	"ip2long":    Ip2long,
	"long2ip":    Long2ip,
	"decodeHtml": DecodeHtml,
}

func Ip2long(ipstr string) uint32 {
	ip := net.ParseIP(ipstr)
	if ip == nil {
		return 0
	}
	ip = ip.To4()
	if ip == nil {
		return 0
	}

	return binary.BigEndian.Uint32(ip)
}

func Long2ip(ipLong uint32) string {
	ipByte := make([]byte, 4)
	binary.BigEndian.PutUint32(ipByte, ipLong)
	ip := net.IP(ipByte)
	return ip.String()
}

func ShowStatus(ts interface{}, on string, on2 string) (tsf string) {
	tsf = ""
	if ts == nil {
		return tsf
	}

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

func DecodeHtml(str string) interface{} {
	fmt.Println(str)
	html := template.HTML(str)
	return html
}

func UploadFile(c *gin.Context, key string) (path string, err error) {
	// FormFile方法会读取参数“upload”后面的文件名，返回值是一个File指针，和一个FileHeader指针，和一个err错误。
	file, header, err := c.Request.FormFile(key)
	if err != nil {
		return "", errors.New("错误的请求")
	}
	// header调用Filename方法，就可以得到文件名
	filename := header.Filename
	fmt.Println(file, err, filename)

	// 创建一个文件，文件名为filename，这里的返回值out也是一个File指针
	out, err := os.Create(filename)
	if err != nil {
		return "", errors.New("创建文件出错")
	}

	defer out.Close()

	// 将file的内容拷贝到out
	_, err = io.Copy(out, file)
	if err != nil {
		return "", errors.New("拷贝文件出错")
	}

	return filename, nil
}
