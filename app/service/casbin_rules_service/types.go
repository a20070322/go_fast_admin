package casbin_rules_service

import (
	"github.com/a20070322/go_fast_admin/ent"
	"github.com/a20070322/go_fast_admin/utils/ent_utils"
)

type FormList struct {
	ent_utils.PageOptions
}

type RepList struct {
	Data  []*ent.CasbinRules `json:"data"`
	Total int              `json:"total"`
}

type FormCreate struct {Ptype string `json:"ptype" binding:"required"`
  V0 string `json:"v0" binding:"required"`
  V1 string `json:"v1" binding:"required"`
  V2 string `json:"v2" binding:"required"`
  V3 string `json:"v3" binding:"required"`
  V4 string `json:"v4" binding:"required"`
  V5 string `json:"v5" binding:"required"`
  }

type RepCreate struct {
	*ent.CasbinRules
}

type FormUpdate struct {Ptype string `json:"ptype" binding:"required"`
  V0 string `json:"v0" binding:"required"`
  V1 string `json:"v1" binding:"required"`
  V2 string `json:"v2" binding:"required"`
  V3 string `json:"v3" binding:"required"`
  V4 string `json:"v4" binding:"required"`
  V5 string `json:"v5" binding:"required"`
  }

type RepUpdate struct {
	*ent.CasbinRules
}