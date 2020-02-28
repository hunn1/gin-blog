package auth

import (
	"Kronos/app/controllers/admin"
	"Kronos/app/models"
	"Kronos/helpers"
	"Kronos/library/apgs"
	"Kronos/library/databases"
	"Kronos/library/password"
	"Kronos/library/session"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginHandler struct {
	admin.AdminBaseHandler
}

// 登录页面显示
func (l LoginHandler) ShowLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "admin_login.html", nil)
}

// 登录
func (l LoginHandler) Login(c *gin.Context) {
	username, pass := c.PostForm("username"), c.PostForm("password")
	var admin models.Admin
	adminData := databases.DB.Where("username=?", username).First(&admin)
	if adminData.Error != nil {
		c.JSON(200, apgs.NewApiReturn(400, "账号或密码错误", nil))
		return
	}
	passBool := password.Compare(admin.Password, pass)
	if passBool != nil {
		c.JSON(200, apgs.NewApiReturn(400, "账号或密码错误", nil))
		return
	}
	admin.Password = ""
	session.SaveSession(c, session.UserKey, admin)

	v := l.GetWhere(1)
	ip := c.ClientIP()
	fmt.Println(helpers.Ip2long(ip))

	v["last_login_ip"] = ip
	databases.DB.Model(&adminData).Where("id = ?", admin.ID).Save(v)
	c.Redirect(302, "/admin/")
}

// 登出
func (l LoginHandler) Logout(c *gin.Context) {
	if hasSession := session.HadSession(c); hasSession == false {
		//c.JSON(200, apgs.NewApiReturn(200, "未进行登录", nil))
		return
	}
	session.ClearAuthSession(c)
	//c.JSON(200, helpers.NewApiReturn(200, "退出成功", nil))
	c.Redirect(302, "/admin/login")
}
