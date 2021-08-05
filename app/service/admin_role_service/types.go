package admin_role_service

import (
	"github.com/a20070322/go_fast_admin/ent"
	"github.com/a20070322/go_fast_admin/utils/ent_utils"
)

type FormList struct {
	ent_utils.PageOptions
	Name string `json:"name" form:"name"`
	IsEnable string  `json:"is_enable" form:"is_enable"`
	IsAll string `json:"is_all" form:"is_all"`
}

type RepList struct {
	Data  []*ent.AdminRole `json:"data"`
	Total int              `json:"total"`
}

type FormCreate struct {
	Name     string `json:"name" binding:"required"`
	IsEnable bool   `json:"is_enable"`
}

type RepCreate struct {
	*ent.AdminRole
}

type FormUpdate struct {
	Name     string `json:"name" binding:"required"`
	IsEnable bool   `json:"is_enable"`
}

type RepUpdate struct {
	*ent.AdminRole
}


type FormSetMenus struct {
	Menus []int `json:"menus"`
}