package admin_controller

import (
	"github.com/a20070322/go_fast_admin/app/service/admin_dict_key_service"
	"github.com/a20070322/go_fast_admin/utils/response"
	"github.com/gin-gonic/gin"
	"net/http"
       "strconv"
    )

type AdminDictKey struct {

}

func (c AdminDictKey) List(ctx *gin.Context) {
	var form admin_dict_key_service.FormList
	if err := ctx.ShouldBind(&form); err != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}
	rep, err := admin_dict_key_service.Init(ctx).List(&form)
	if err != nil {
		response.Fail(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	response.Success(ctx, "ok", rep)
}

func (c AdminDictKey) Create(ctx *gin.Context) {
	var form admin_dict_key_service.FormCreate
	if err := ctx.ShouldBind(&form); err != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}
	rep, err := admin_dict_key_service.Init(ctx).Create(&form)
	if err != nil {
		response.Fail(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	response.Success(ctx, "ok", rep)
}

func (c AdminDictKey) Update(ctx *gin.Context) {
        id, err := strconv.Atoi(ctx.Param("id"))
        if err != nil {
            response.Fail(ctx, http.StatusUnprocessableEntity, "参数错误", nil)
            return
        }
    var form admin_dict_key_service.FormUpdate
	if err := ctx.ShouldBind(&form); err != nil {
		response.Fail(ctx, http.StatusUnprocessableEntity, err.Error(), nil)
		return
	}
	rep, err := admin_dict_key_service.Init(ctx).Update(id, &form)
	if err != nil {
		response.Fail(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	response.Success(ctx, "ok", rep)
}

func (c AdminDictKey) Delete (ctx *gin.Context) {
    var err error
        id, err := strconv.Atoi(ctx.Param("id"))
        if err != nil {
            response.Fail(ctx, http.StatusUnprocessableEntity, "参数错误", nil)
            return
        }
    err = admin_dict_key_service.Init(ctx).Delete(id)
	if err != nil {
		response.Fail(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	response.Success(ctx, "ok", "")
}
