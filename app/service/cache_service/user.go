package cache_service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/a20070322/go_fast_admin/app/service/admin_user_service"
	"github.com/a20070322/go_fast_admin/ent"
	"github.com/a20070322/go_fast_admin/global"
	"time"
)

//判断用户是否被缓存
func (m Cache) CheckAdminUserCatch(id string) bool {
	b,err := m.Rdb.Exists(m.ctx,fmt.Sprintf("%s%s", m.AdminUserPrefix, id)).Result()
	if err != nil {
		global.Logger.Error(err)
	}
	return b>0
}
//设置用户缓存
func (m Cache) SetAdminUserCatch(u *ent.AdminUser) error {
	if u == nil {
		return errors.New("ent.AdminUser is nil")
	}
	str, err := json.Marshal(u)
	if err != nil {
		global.Logger.Error(err)
		return err
	}
	rep := m.Rdb.Set(m.ctx, fmt.Sprintf("%s%s", m.AdminUserPrefix, u.ID.String()), string(str), time.Second*time.Duration(global.AppSetting.Jwt.TokenExpireDuration))
	if rep.Err() != nil {
		global.Logger.Error(rep.Err())
		return rep.Err()
	}
	return nil
}

//读取用户缓存
func (m Cache) GetAdminUserCatch(id string) (*ent.AdminUser, error) {
	if id == "" {
		return nil, errors.New("id is empty")
	}
	rep, err := m.Rdb.Get(m.ctx, fmt.Sprintf("%s%s", m.AdminUserPrefix, id)).Result()
	if err != nil {
		global.Logger.Error(err)
		u, err := admin_user_service.Init(m.ctx).FindById(id)
		_ = m.SetAdminUserCatch(u)
		if err != nil {
			return nil, err
		}
		return u, err
	}
	var u ent.AdminUser
	err = json.Unmarshal([]byte(rep), &u)
	if err != nil {
		global.Logger.Error(err)
		return nil, err
	}
	return &u, nil
}