package router

import (
	"github.com/a20070322/go_fast_admin/app/controller/admin_controller"
	"github.com/a20070322/go_fast_admin/middleware"
	"github.com/a20070322/go_fast_admin/utils/jwt"
	"github.com/gin-gonic/gin"
)

func AdminRouter(r *gin.RouterGroup) {
	//.Use(middleware.JwtAuth(jwt.UserGroupAdmin))
	api := r.Group("/admin")

	//公共模块不做权限验证
	common:= api.Group("/common").Use(middleware.JwtAuth(jwt.UserGroupAdmin))
	common.GET("/oneself_menus", admin_controller.AdminMenus{}.GetUserMenu)

	//auth模块
	auth := api.Group("/auth")
	auth.POST("/login", admin_controller.AdminAuth{}.Login)
	auth.POST("/refresh_token", admin_controller.AdminAuth{}.RefreshToken)

	//user模块
	user := api.Group("/user").Use(middleware.JwtAuth(jwt.UserGroupAdmin),middleware.AdminRbac())
	user.GET("/list", admin_controller.AdminUser{}.List)
	user.POST("/create", admin_controller.AdminUser{}.Create)
	user.POST("/update/:id", admin_controller.AdminUser{}.Update)
	user.POST("/delete/:id", admin_controller.AdminUser{}.Delete)


	//AdminMenus模块
	AdminMenus := api.Group("/admin_menus").Use(middleware.JwtAuth(jwt.UserGroupAdmin),middleware.AdminRbac())
	//AdminMenus.GET("/list", admin_controller.AdminMenus{}.List)
	AdminMenus.GET("/tree_list", admin_controller.AdminMenus{}.TreeList)
	AdminMenus.POST("/create", admin_controller.AdminMenus{}.Create)
	AdminMenus.POST("/update/:id", admin_controller.AdminMenus{}.Update)
	AdminMenus.POST("/delete/:id", admin_controller.AdminMenus{}.Delete)


	//AdminRole模块
	AdminRole := api.Group("/admin_role").Use(middleware.JwtAuth(jwt.UserGroupAdmin),middleware.AdminRbac())
	AdminRole.GET("/list", admin_controller.AdminRole{}.List)
	AdminRole.POST("/create", admin_controller.AdminRole{}.Create)
	AdminRole.POST("/update/:id", admin_controller.AdminRole{}.Update)
	AdminRole.POST("/delete/:id", admin_controller.AdminRole{}.Delete)
	AdminRole.GET("/find_menus/:id", admin_controller.AdminRole{}.FindMenus)
	AdminRole.POST("/menu_update/:id", admin_controller.AdminRole{}.SetMenus)
	//{{.RouteTmpl}}
}
