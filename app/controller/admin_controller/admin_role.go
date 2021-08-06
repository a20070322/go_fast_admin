package admin_controller

import (
	"github.com/a20070322/go_fast_admin/app/service/admin_role_service"
	"github.com/a20070322/go_fast_admin/app/service/cache_service"
	"github.com/a20070322/go_fast_admin/global"
	"github.com/a20070322/go_fast_admin/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AdminRole struct {
}

func (c AdminRole) List(ctx *gin.Context) {
	var form admin_role_service.FormList
	if err := ctx.ShouldBind(&form); err != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}
	rep, err := admin_role_service.Init(ctx).List(&form)
	if err != nil {
		response.Fail(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	response.Success(ctx, "ok", rep)
}

func (c AdminRole) Create(ctx *gin.Context) {
	var form admin_role_service.FormCreate
	if err := ctx.ShouldBind(&form); err != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}
	rep, err := admin_role_service.Init(ctx).Create(&form)
	if err != nil {
		response.Fail(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	response.Success(ctx, "ok", rep)
}

func (c AdminRole) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, "参数错误", nil)
		return
	}
	var form admin_role_service.FormUpdate
	if err := ctx.ShouldBind(&form); err != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}
	rep, err := admin_role_service.Init(ctx).Update(id, &form)
	if err != nil {
		response.Fail(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	catch:=  cache_service.Init(ctx)
	if catch.CheckAdminRoleCatch(id) {
		u, err := admin_role_service.Init(ctx).FindByIdWithMenu(id)
		if err != nil {
			global.Logger.Error(err)
		}
		err = catch.SetAdminRoleCatch(u)
		if err != nil {
			global.Logger.Error(err)
		}
	}
	response.Success(ctx, "ok", rep)
}

func (c AdminRole) Delete(ctx *gin.Context) {
	var err error
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, "参数错误", nil)
		return
	}
	err = admin_role_service.Init(ctx).Delete(id)
	if err != nil {
		response.Fail(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	response.Success(ctx, "ok", "")
}

func (c AdminRole) FindMenus(ctx *gin.Context) {
	var err error
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, "参数错误", nil)
		return
	}
	rep, err := admin_role_service.Init(ctx).FindMenus(id)
	if err != nil {
		response.Fail(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	response.Success(ctx, "ok", rep)
}

func (c AdminRole) SetMenus(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, "参数错误", nil)
	}
	var form admin_role_service.FormSetMenus
	if err := ctx.ShouldBind(&form); err != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}
	err = admin_role_service.Init(ctx).SetMenus(id, &form)
	catch:=  cache_service.Init(ctx)
	if catch.CheckAdminRoleCatch(id) {
		u, err := admin_role_service.Init(ctx).FindByIdWithMenu(id)
		if err != nil {
			global.Logger.Error(err)
		}
		err = catch.SetAdminRoleCatch(u)
		if err != nil {
			global.Logger.Error(err)
		}
	}
	if err != nil {
		response.Fail(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	response.Success(ctx, "ok", "")
}
