package admin_dict_service

import (
	"github.com/a20070322/go_fast_admin/ent"
	"github.com/a20070322/go_fast_admin/utils/ent_utils"
)

type FormList struct {
	ent_utils.PageOptions
	DictName string `json:"dict_name" form:"dict_name"`
	DictType string `json:"dict_type" form:"dict_type"`
	IsEnable string `json:"is_enable" form:"is_enable"`
}

type RepList struct {
	Data  []*ent.AdminDict `json:"data"`
	Total int              `json:"total"`
}

type FormCreate struct {
	DictType string `json:"dict_type" binding:"required"`
	DictName string `json:"dict_name" binding:"required"`
	Remarks  string `json:"remarks"`
	IsEnable bool   `json:"is_enable" binding:"required"`
}

type RepCreate struct {
	*ent.AdminDict
}

type FormUpdate struct {
	DictType string `json:"dict_type" binding:"required"`
	DictName string `json:"dict_name" binding:"required"`
	Remarks  string `json:"remarks"`
	IsEnable bool   `json:"is_enable" binding:"required"`
}

type RepUpdate struct {
	*ent.AdminDict
}

type DictMap struct {
	Label     string `json:"label"`
	Value     string `json:"value"`
}
