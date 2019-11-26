package boot

import (
	"Kronos/library/databases"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"time"
)

/**
启动框架
*/
func Run(router *gin.Engine) {

	server := &http.Server{
		Addr:    viper.GetString("addr"),
		Handler: router,
	}
	// 开启 Server
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	// 数据库初始化
	dbType := viper.GetString("db.type")
	host := viper.GetString("db.host")

	user := viper.GetString("db.user")
	pass := viper.GetString("db.pass")
	dbname := viper.GetString("db.dbname")
	charset := viper.GetString("db.charset")
	loc := viper.GetString("db.loc")
	native := viper.GetString("db.native")
	databases.InitDB(dbType, host, user, pass, dbname, charset, loc, url.QueryEscape(native))
	// 接收退出信号
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server")
	// 超时处理
	timeout, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancelFunc()

	if err := server.Shutdown(timeout); err != nil {

		log.Fatal("Server Shutdown", err)
	}

	log.Println("Server exiting")
	// PID 文件处理
	pid := fmt.Sprintf("%d", os.Getpid())

	_, openErr := os.OpenFile("pid", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if openErr == nil {
		_ = ioutil.WriteFile("pid", []byte(pid), 0)
	}
}
