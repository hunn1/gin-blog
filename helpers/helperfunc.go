package helpers

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"
	"html/template"
	"mime/multipart"
	"net"
	"net/http"
	"os"
	"path"
	"time"
)

var Builtins = template.FuncMap{
	"showStatus":    ShowStatus,
	"showtime":      ShowTime,
	"ip2long":       Ip2long,
	"long2ip":       Long2ip,
	"decodeHtml":    DecodeHtml,
	"getUploadPath": GetUploadPath,
}

const UploadPath = "./resources/public/thumb/"
const UploadUrl = "/public/thumb/"

func GetUploadPath(filename string) string {

	if filename == "" {
		return UploadUrl + "none.jpg"
	}
	return UploadUrl + filename
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

func GetFileContentType(out multipart.File) (string, error) {
	// 只需要前 512 个字节就可以了
	buffer := make([]byte, 512)
	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}
	contentType := http.DetectContentType(buffer)
	return contentType, nil
}

func UploadFile(c *gin.Context, key string) (filename string, err error) {
	// FormFile方法会读取参数“upload”后面的文件名，返回值是一个File指针，和一个FileHeader指针，和一个err错误。
	header, err := c.FormFile(key)
	if err == nil {
		err = CheckDir(UploadPath, 0711)
		if err != nil {
			return "", errors.New("文件夹不存在无法创建")
		}
		k := ksuid.New()
		open, err := header.Open()

		contentType, err := GetFileContentType(open)
		isSupportFile := false
		allowTypes := []string{
			"image/png",
			"image/jpg",
			"image/jpeg",
			"image/gif",
			"image/webp",
		}
		if len(allowTypes) != 0 {
			for i := 0; i < len(allowTypes); i++ {
				if allowTypes[i] == contentType {
					isSupportFile = true
					break
				}
			}
			if isSupportFile == false {
				return "", errors.New("不支持的文件类型")
			}
		}
		filename = k.String() + path.Ext(header.Filename)
		err = c.SaveUploadedFile(header, UploadPath+filename)
		if err != nil {
			return "", err
		}
		return filename, nil
	}
	return "", nil

}

func CheckDir(path string, perm os.FileMode) error {
	// check
	if _, err := os.Stat(path); err == nil {
		return nil
	} else {
		err := os.MkdirAll(path, perm)
		if err != nil {
			return err
		}
	}

	// check again
	if _, err := os.Stat(path); err == nil {
		return nil
	}
	return nil
}
