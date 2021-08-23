package admin_controller

import (
	"github.com/a20070322/go_fast_admin/app/service/admin_auth_service"
	"github.com/a20070322/go_fast_admin/app/service/admin_user_service"
	"github.com/a20070322/go_fast_admin/utils/jwt"
	"github.com/a20070322/go_fast_admin/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AdminAuth struct {
}

func (c AdminAuth) Login(ctx *gin.Context) {
	var form admin_auth_service.FormLogin
	if err := ctx.ShouldBind(&form); err != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}
	rep, err := admin_auth_service.Init(ctx).Login(&form)
	if err != nil {
		response.Fail(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	response.Success(ctx, "ok", rep)
}

type FormRefreshToken struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// 刷新token
func (c AdminAuth) RefreshToken(ctx *gin.Context) {
	var form FormRefreshToken
	if err := ctx.ShouldBind(&form); err != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}
	claims, err2 := jwt.VerifyAction(form.RefreshToken)
	if err2 != nil || claims.IsRefreshToken != true {
		response.Fail(ctx, http.StatusUnprocessableEntity, "refresh_token格式错误", nil)
		return
	}
	isCheck, err3 := jwt.CheckTokenCatch(claims, form.RefreshToken, true)
	if err3 != nil {
		response.Fail(ctx, http.StatusBadGateway, "服务器内部错误", nil)
		return
	}
	if isCheck != true {
		response.Fail(ctx, http.StatusUnprocessableEntity, "refresh_token已失效", nil)
		return
	}
	user, err3 := admin_user_service.Init(ctx).FindById(claims.UserID)
	if err3 != nil && user != nil {
		response.Fail(ctx, http.StatusInternalServerError, err2.Error(), nil)
		return
	}
	rep, err4 := admin_auth_service.Init(ctx).GetToken(user)
	if err4 != nil {
		response.Fail(ctx, http.StatusInternalServerError, err4.Error(), nil)
		return
	}
	response.Success(ctx, "ok", rep)
}
