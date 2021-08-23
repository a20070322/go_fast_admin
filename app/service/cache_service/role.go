package cache_service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/a20070322/go_fast_admin/app/service/admin_role_service"
	"github.com/a20070322/go_fast_admin/ent"
	"github.com/a20070322/go_fast_admin/global"
)

//设置角色缓存
func (m Cache) SetAdminRoleCatch(u *ent.AdminRole) error {
	if u == nil {
		return errors.New("ent.AdminRole is nil")
	}
	global.Logger.Debug("角色缓存更新")
	str, err := json.Marshal(u)
	if err != nil {
		global.Logger.Error(err)
		return err
	}
	rep := m.Rdb.Set(m.ctx, fmt.Sprintf("%s%d", m.AdminRolePrefix, u.ID), string(str),0)
	if rep.Err() != nil {
		global.Logger.Error(rep.Err())
		return rep.Err()
	}
	return nil
}

//判断角色是否被缓存
func (m Cache) CheckAdminRoleCatch(id int) bool {
	b,err := m.Rdb.Exists(m.ctx,fmt.Sprintf("%s%d", m.AdminRolePrefix, id)).Result()
	if err != nil {
		global.Logger.Error(err)
	}
	return b>0
}

//读取角色缓存
func (m Cache) GetAdminRoleCatch(id int) (*ent.AdminRole, error) {
	if id == 0 {
		return nil, errors.New("id is 0")
	}
	rep, err := m.Rdb.Get(m.ctx, fmt.Sprintf("%s%d", m.AdminRolePrefix, id)).Result()
	if err != nil {
		global.Logger.Error(err)
		u, err := admin_role_service.Init(m.ctx).FindByIdWithMenu(id)
		_ = m.SetAdminRoleCatch(u)
		if err != nil {
			global.Logger.Error(err)
			return nil, err
		}
		return u, err
	}
	var u ent.AdminRole
	err = json.Unmarshal([]byte(rep), &u)
	if err != nil {
		global.Logger.Error(err)
		return nil, err
	}
	return &u, nil
}