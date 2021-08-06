package middleware

import (
	"github.com/a20070322/go_fast_admin/app/service/check_role_service"
	"github.com/a20070322/go_fast_admin/global"
	"github.com/a20070322/go_fast_admin/utils/response"
	"github.com/gin-gonic/gin"
)

//目前只针对后台管理系统使用
func AdminRbac() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid := c.GetString("uid")
		if uid == "" {
			response.NoAuth(c, 403, "无效token", "50001")
		}
		b, err := check_role_service.AdminCheckRole(c, &check_role_service.FormAdminCheckRole{
			Uid:    uid,
			Path:   c.Request.URL.Path,
			Method: c.Request.Method,
		})
		if err != nil {
			global.Logger.Error(err)
			response.NoAuth(c, 500, "服务器内部错误", "50002")
		}
		if b {
			c.Next()
		}else{
			response.NoAuth(c, 403, "无权限访问", "50003")
		}
	}
}
