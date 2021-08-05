package admin_controller

import (
	"fmt"
	"github.com/a20070322/go_fast_admin/app/service/admin_menus_service"
	"github.com/a20070322/go_fast_admin/app/service/admin_user_service"
	"github.com/a20070322/go_fast_admin/utils/jwt"
	"github.com/a20070322/go_fast_admin/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AdminMenus struct {
}

func (c AdminMenus) List(ctx *gin.Context) {
	var form admin_menus_service.FormList
	if err := ctx.ShouldBind(&form); err != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}
	rep, err := admin_menus_service.Init(ctx).List(&form)
	if err != nil {
		response.Fail(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	response.Success(ctx, "ok", rep)
}

func (c AdminMenus) TreeList(ctx *gin.Context) {
	var form admin_menus_service.FormList
	if err := ctx.ShouldBind(&form); err != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}
	rep, err := admin_menus_service.Init(ctx).TreeList(&form)
	if err != nil {
		response.Fail(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	response.Success(ctx, "ok", rep)
}

func (c AdminMenus) Create(ctx *gin.Context) {
	var form admin_menus_service.FormCreate
	if err := ctx.ShouldBind(&form); err != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}
	rep, err := admin_menus_service.Init(ctx).Create(&form)
	if err != nil {
		response.Fail(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	response.Success(ctx, "ok", rep)
}

func (c AdminMenus) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, "参数错误", nil)
	}
	var form admin_menus_service.FormUpdate
	if err := ctx.ShouldBind(&form); err != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}
	rep, err := admin_menus_service.Init(ctx).Update(id, &form)
	if err != nil {
		response.Fail(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	response.Success(ctx, "ok", rep)
}

func (c AdminMenus) Delete(ctx *gin.Context) {
	var err error
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, "参数错误", nil)
	}
	err = admin_menus_service.Init(ctx).Delete(id)
	if err != nil {
		response.Fail(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	response.Success(ctx, "ok", "")
}

func (c AdminMenus) GetUserMenu(ctx *gin.Context) {
	uid, err := jwt.GetTokenId(ctx)
	if err != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, err.Error(), nil)
	}
	fmt.Println(uid,"123")
	if uid == ""{
		response.Fail(ctx, http.StatusUnprocessableEntity, "用户未找到", nil)
	}
	user, err := admin_user_service.Init(ctx).FindById(uid)
	//用户角色列表
	var roles []int
	if err != nil {
		response.Fail(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	for _, v := range user.Edges.Role {
		roles = append(roles, v.ID)
	}

	rep,err := admin_menus_service.Init(ctx).GetUserMenu(roles)
	if err != nil {
		response.Fail(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	response.Success(ctx, "ok", rep)
}


