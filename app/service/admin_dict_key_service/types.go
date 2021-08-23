package admin_dict_key_service

import (
	"github.com/a20070322/go_fast_admin/ent"
	"github.com/a20070322/go_fast_admin/utils/ent_utils"
)

type FormList struct {
	ent_utils.PageOptions
	Fid int `json:"fid" form:"fid""`
}

type RepList struct {
	Data  []*ent.AdminDictKey `json:"data"`
	Total int                 `json:"total"`
}

type FormCreate struct {
	DictLabel string `json:"dict_label" binding:"required"`
	DictCode  string `json:"dict_code" binding:"required"`
	Fid       int    `json:"fid" binding:"required"`
	Sort      int    `json:"sort"`
	Remarks   string `json:"remarks"`
	IsEnable  bool   `json:"is_enable"`
}

type RepCreate struct {
	*ent.AdminDictKey
}

type FormUpdate struct {
	DictLabel string `json:"dict_label" binding:"required"`
	DictCode  string `json:"dict_code" binding:"required"`
	Fid       int    `json:"fid" binding:"required"`
	Sort      int    `json:"sort"`
	Remarks   string `json:"remarks"`
	IsEnable  bool   `json:"is_enable"`
}

type RepUpdate struct {
	*ent.AdminDictKey
}
