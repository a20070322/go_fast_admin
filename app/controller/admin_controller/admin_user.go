package admin_controller

import (
	"github.com/a20070322/go_fast_admin/app/service/admin_user_service"
	"github.com/a20070322/go_fast_admin/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AdminUser struct {
}


func (c AdminUser) List(ctx *gin.Context) {
	var form admin_user_service.FromList
	if err := ctx.ShouldBind(&form); err != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}
	rep, err := admin_user_service.Init(ctx).List(&form)
	if err != nil {
		response.Fail(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	response.Success(ctx, "ok", rep)
}

func (c AdminUser) Create(ctx *gin.Context) {
	var form admin_user_service.FromCommon
	if err := ctx.ShouldBind(&form); err != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}
	form.Password="fast_admin"
	rep, err := admin_user_service.Init(ctx).Create(&form)
	if err != nil {
		response.Fail(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	response.Success(ctx, "ok", rep)
}

func (c AdminUser) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		response.Fail(ctx, http.StatusUnprocessableEntity, "参数错误", nil)
		return
	}
	var form admin_user_service.FromUpdate
	if err := ctx.ShouldBind(&form); err != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}
	rep, err := admin_user_service.Init(ctx).Update(id, &form)
	if err != nil {
		response.Fail(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	response.Success(ctx, "ok", rep)
}

func (c AdminUser) Delete (ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		response.Fail(ctx, http.StatusUnprocessableEntity, "参数错误", nil)
		return
	}
	err := admin_user_service.Init(ctx).Delete(id)
	if err != nil {
		response.Fail(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	response.Success(ctx, "ok", "")
}
