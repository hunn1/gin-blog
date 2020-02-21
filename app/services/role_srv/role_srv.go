package role_srv

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func ApplyRole(c *gin.Context) {
	id, _ := strconv.ParseInt(c.DefaultPostForm("id", "0"), 10, 0)
	permissions_id, _ := strconv.ParseInt(c.DefaultPostForm("permissions_id", "0"), 10, 0)
	fmt.Println(id, permissions_id)
}
