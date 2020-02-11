package admin

import (
	"Kronos/helpers"
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
	c.JSON(200, helpers.ApiReturn{
		Code:    0,
		Message: "111",
		Data:    nil,
	})
}
func Login(c *gin.Context) {
	//username, password := c.PostForm("username"), c.PostForm("password")
	//// Authentication
	//// blahblah...
	//
	//// Generate random session id
	//u, err := uuid.NewRandom()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//sessionId := fmt.Sprintf("%s-%s", u.String(), username)
	//// Store current subject in cache
	//component.GlobalCache.Set(sessionId, []byte(username))
	//// Send cache key back to client in cookie
	//c.SetCookie("current_subject", sessionId, 30*60, "/resource", "", false, true)
	//c.JSON(200, component.RestResponse{Code: 1, Message: username + " logged in successfully"})
}
