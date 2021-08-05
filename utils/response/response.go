package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ReturnJson(context *gin.Context, httpCode int, dataCode int, msg string, data interface{}) {
	resp := gin.H{
		"code": dataCode,
		"msg":  msg,
		"data": data,
	}
	//response, _ := json.Marshal(resp)
	context.Set("response", resp)
	context.JSON(httpCode,resp)
}

// 成功返回
func Success(c *gin.Context, msg string, data interface{}) {
	ReturnJson(c, http.StatusOK, http.StatusOK, msg, data)
}

// 失败返回
func Fail(c *gin.Context, dataCode int, msg string, data interface{}) {
	ReturnJson(c, http.StatusBadRequest, dataCode, msg, data)
	c.Abort()
}

// 权限失败返回
func NoAuth(c *gin.Context, dataCode int, msg string, data interface{}) {
	ReturnJson(c, http.StatusUnauthorized, dataCode, msg, data)
	c.Abort()
}


// 权限失败返回
func NoAuth403(c *gin.Context, dataCode int, msg string, data interface{}) {
	ReturnJson(c, http.StatusForbidden, dataCode, msg, data)
	c.Abort()
}