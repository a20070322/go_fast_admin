package middleware

import (
	"github.com/gin-gonic/gin"
)

//目前只针对后台管理系统使用
func Casbin() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		//uid := c.GetString("uid")
		//if uid == "" {
		//	response.NoAuth(c, 403, "无效token", "50001")
		//}
		//res, err := global.Rbac.Enforce(uid,c.Request.URL.Path, c.Request.Method)
		//if err != nil {
		//	response.Fail(c,http.StatusBadGateway,"服务器内部错误","50002")
		//}
		//if res {
		//	c.Next()
		//} else {
		//	// deny the request, show an error
		//	response.Fail(c,http.StatusForbidden,"无权访问","50003")
		//}
	}
}
