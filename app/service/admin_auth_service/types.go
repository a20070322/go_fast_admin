package admin_auth_service

import (
	"github.com/a20070322/go_fast_admin/ent"
	"github.com/a20070322/go_fast_admin/utils/jwt"
)

//login form
type FormLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//返回token
type RepGetToken struct {
	JwtData *jwt.AuthReturn `json:"jwt_data"`
	User    *ent.AdminUser   `json:"user"`
}