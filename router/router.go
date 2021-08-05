package router

import (
	"context"
	"fmt"
	"github.com/a20070322/go_fast_admin/global"
	"github.com/a20070322/go_fast_admin/middleware"
	"github.com/a20070322/go_fast_admin/types"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/color"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"
)

func StartServer() {
	router := gin.New()
	// 本地开发引入自带中间件比较美观
	if global.AppSetting.Env == types.EnvLOCAL {
		router.Use(gin.Logger(), gin.Recovery())
	}
	if global.AppSetting.Env != types.EnvLOCAL {
		router.Use(middleware.Recovery(true))
		// 日志记录至文件
		router.Use(middleware.Logger())
	}
	router.GET("/", func(context *gin.Context) {
		context.String(200, "ok")
	})
	// api注入
	api := router.Group("api")
	apiRegisterApis(api)
	fmt.Println(fmt.Sprintf("ServerPort: %s", color.Green(global.AppSetting.Server.Port)))
	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(global.AppSetting.Server.Port),
		Handler: router,
	}
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	global.Db.Close()
	global.Rdb.Close()
	log.Println("Server exiting")
}

func apiRegisterApis(api *gin.RouterGroup) {
	AdminRouter(api)
}