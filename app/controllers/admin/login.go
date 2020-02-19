package admin

import (
	"Kronos/app/models"
	"Kronos/helpers"
	"Kronos/library/databases"
	"Kronos/library/password"
	"Kronos/library/session"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 登录页面显示
func ShowLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "admin_login.html", nil)
}

// 登录
func Login(c *gin.Context) {
	username, pass := c.PostForm("username"), c.PostForm("password")
	var admin models.Admin
	adminData := databases.DB.Where("username=?", username).First(&admin)
	if adminData.Error != nil {
		c.JSON(200, helpers.NewApiReturn(400, "账号或密码错误", nil))
		return
	}
	passBool := password.Compare(admin.Password, pass)
	if passBool != nil {
		c.JSON(200, helpers.NewApiReturn(400, "账号或密码错误", nil))
		return
	}
	session.SaveSession(c, uint(admin.ID))
	c.Redirect(302, "/admin/")
}

// 登出
func Logout(c *gin.Context) {
	if hasSession := session.HadSession(c); hasSession == false {
		c.JSON(200, helpers.NewApiReturn(200, "未进行登录", nil))
		return
	}
	session.ClearAuthSession(c)
	c.JSON(200, helpers.NewApiReturn(200, "退出成功", nil))
}
