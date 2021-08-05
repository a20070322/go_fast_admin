package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func PriceToStr(price int) string {
	str := strconv.Itoa(price)
	if len(str) < 3 {
		for num := 3-len(str);num>0;num--  {
			str = "0"+str
		}
	}
	str2 := string([]rune(str)[:len(str)-2])
	str3 := string([]rune(str)[len(str)-2 : len(str)])
	return str2 + "." + str3
}


func GetTracingID(ctx *gin.Context) string {
	if tracingID, tracingIDBool := ctx.Get("TracingID"); tracingIDBool == true {
		return tracingID.(string)
	} else {
		return "no tracingID"
	}
}