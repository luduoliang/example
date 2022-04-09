package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//验证后台用户是否登录
func CheckAdminLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		admin_info := GetSession(c, "admin_info")
		if admin_info == nil {
			c.Redirect(http.StatusFound, "/admin/login")
			return
		}

		c.Next()
	}
}
