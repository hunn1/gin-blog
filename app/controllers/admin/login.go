package admin

import (
	"Kronos/helpers"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowLogin(c *gin.Context) {
	//session := sessions.Default(c)
	//
	//session.Set("loginuser", "Testa")
	//session.Save()
	//
	//loginuser := session.Get("loginuser")
	//fmt.Println("loginuser:", loginuser)

	c.HTML(http.StatusOK, "admin_login.html", nil)
}

func TestC(c *gin.Context) {
	ginview.HTML(c, http.StatusUnauthorized, "err/401", helpers.NewApiRedirect(200, "无权限访问该内容", "/admin/login"))
	c.Abort()
	//c.JSON(200, helpers.NewApiReturn(0, "111", nil))
}

func Login(c *gin.Context) {
	//session := sessions.Default(c)

	//username, password := c.PostForm("username"), c.PostForm("password")

	//c.JSON(200, helpers.NewApiReturn(200, "", gin.H{"username": username, "password": session.Get("Test")}))
	//// Authentication
	//// blahblah...
	//
	//// Generate random session id
	//u, err := .NewRandom()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//sessionId := fmt.Sprintf("%s-%s", u.String(), username)
	//// Store current subject in cache
	//component.GlobalCache.Set(sessionId, []byte(username))
	//// Send cache key back to client in cookie
	//c.SetCookie("current_subject", sessionId, 30*60, "/resource", "", false, true)
	//c.JSON(200, helpers.ApiReturn{ 200, username + " logged in successfully", nil})
}
