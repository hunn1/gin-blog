package logs

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotators "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path"
	"time"
)

// 日志记录到文件
func LoggerToFile() gin.HandlerFunc {

	pwd, _ := os.Getwd()

	logFilePath := pwd + "/" + viper.GetString("locus.path")

	logFileName := viper.GetString("locus.file")

	// 日志文件
	fileName := path.Join(logFilePath, logFileName)

	// 写入文件
	src, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		fmt.Println("err", err)
	}

	// 实例化
	logger := logrus.New()

	// 设置输出
	logger.Out = src

	// 设置日志级别
	logger.SetLevel(logrus.DebugLevel)

	// 设置
	logWriter, err := rotators.New(
		// 分割后的文件名称
		fileName+".%Y%m%d.log",

		// 生成软链，指向最新日志文件
		rotators.WithLinkName(fileName),

		// 设置最大保存时间(7天)
		rotators.WithMaxAge(7*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotators.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{})

	// 新增钩子
	logger.AddHook(lfHook)

	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		// 日志格式
		logger.WithFields(logrus.Fields{
			"statusCode":  statusCode,
			"latencyTime": latencyTime,
			"clientIP":    clientIP,
			"reqMethod":   reqMethod,
			"reqUri":      reqUri,
		}).Info()
	}
}
