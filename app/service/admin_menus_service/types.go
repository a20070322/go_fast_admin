package admin_menus_service

import (
	"github.com/a20070322/go_fast_admin/ent"
	"github.com/a20070322/go_fast_admin/utils/ent_utils"
)

type FormList struct {
	ent_utils.PageOptions
}

type RepList struct {
	Data  []*ent.AdminMenus `json:"data"`
	Total int               `json:"total"`
}


type FormCreate struct {
	Name           string `json:"name" binding:"required"`
	Path           string `json:"path"`
	RouterPath     string `json:"router_path"`
	Icon           string `json:"icon"`
	Type           int8   `json:"type"`
	PowerStr       string `json:"power_str"`
	Sort           int    `json:"sort"`
	Fid            int    `json:"fid"`
	IsExternalLink bool   `json:"is_external_link"`
	IsShow         bool   `json:"is_show"`
	IsEnable       bool   `json:"is_enable"`
}

type RepCreate struct {
	*ent.AdminMenus
}

type FormUpdate struct {
	Name           string `json:"name" binding:"required"`
	Path           string `json:"path"`
	RouterPath     string `json:"router_path"`
	Icon           string `json:"icon"`
	Type           int8   `json:"type"`
	PowerStr       string `json:"power_str"`
	Sort           int    `json:"sort"`
	Fid            int    `json:"fid"`
	IsExternalLink bool   `json:"is_external_link"`
	IsShow         bool   `json:"is_show"`
	IsEnable       bool   `json:"is_enable"`
}

type RepUpdate struct {
	*ent.AdminMenus
}

type RepGetUserMenu struct {
	Menu []*MenusTree      `json:"menu"`
	Role []*ent.AdminMenus `json:"role"`
}

type MenusTree struct {
	*ent.AdminMenus
	Children []*MenusTree `json:"children"`
}
