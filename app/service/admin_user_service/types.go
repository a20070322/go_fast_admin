package admin_user_service

import (
	"github.com/a20070322/go_fast_admin/ent"
	"github.com/a20070322/go_fast_admin/utils/ent_utils"
)

type FromList struct {
	ent_utils.PageOptions
	Username string `json:"username" form:"username"`
	Phone    string `json:"phone" form:"phone"`
	IsEnable string `json:"is_enable" form:"is_enable"`
	Role     int   `json:"role" form:"role" query:"role"`
}

type RepList struct {
	Data  []*ent.AdminUser `json:"data"`
	Total int              `json:"total"`
}

type FromCommon struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password"`
	Avatar   string `json:"avatar" `
	Phone    string `json:"phone"`
	IsEnable bool   `json:"is_enable"`
	Roles    []int  `json:"roles"`
}

type RepCommon struct {
	*ent.AdminUser
}

type FromUpdate struct {
	Username string `json:"username" binding:"required"`
	//Password string `json:"password" binding:"required"`
	Avatar   string `json:"avatar" `
	Phone    string `json:"phone"`
	IsEnable bool   `json:"is_enable"`
	Roles    []int  `json:"roles"`
}
