package middleware

import (
	"github.com/a20070322/go_fast_admin/utils/jwt"
	"github.com/a20070322/go_fast_admin/utils/response"
	"github.com/gin-gonic/gin"
	"strings"
)

type HeaderParams struct {
	Authorization string `header:"Authorization"`
}

// jwt鉴权中间件
func JwtAuth(group jwt.UserGroupType) gin.HandlerFunc {
	return func(context *gin.Context) {
		headerParams := HeaderParams{}
		if err := context.ShouldBindHeader(&headerParams); err != nil {
			response.NoAuth(context, 401, "token不存在", "40101")
			context.Abort()
			return
		}
		authorization := strings.Split(headerParams.Authorization, " ")
		if len(authorization) == 2 && authorization[0] == "Bearer" {
			claims, err2 := jwt.VerifyAction(authorization[1])
			if err2 != nil {
				response.NoAuth(context, 401, "token验证失败", "40102")
				context.Abort()
				return
			}
			if group != jwt.UserGroupAll && group != claims.UserGroup {
				response.NoAuth(context, 401, "token验证失败", "40110")
				context.Abort()
				return
			}
			if claims.IsRefreshToken == true {
				response.NoAuth(context, 401, "token验证失败", "40111")
				context.Abort()
				return
			}
			boolV, err3 := jwt.CheckTokenCatch(claims, authorization[1], false)
			if err3 != nil {
				response.NoAuth(context, 401, "鉴权失败，重新登录", "50001")
				context.Abort()
				return
			}
			if boolV == false {
				response.NoAuth(context, 401, "已在其他设备登录", "40103")
				context.Abort()
				return
			}
			context.Set("uid", claims.UserID)
			context.Next()
		} else {
			response.NoAuth(context, 401, "token格式错误", "40100")
		}
	}
}
