package middleware

import (
	"fmt"
	"github.com/a20070322/go_fast_admin/global"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-uuid"
	"github.com/labstack/gommon/color"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		tracingID, err := uuid.GenerateUUID()
		if err != nil {
			global.Logger.Error(err.Error())
			tracingID = "no tracingID"
		}
		c.Set("TracingID", tracingID)
		start := time.Now()
		path := c.Request.URL.Path
		c.Next()
		cost := time.Since(start).Milliseconds()
		global.Logger.Info(fmt.Sprintf("[%s]   %dms | %s | %s   \"%s\"", color.Green(c.Writer.Status()), cost, c.ClientIP(), c.Request.Method, path))
	}
}
