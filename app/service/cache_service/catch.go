package cache_service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/a20070322/go_fast_admin/app/service/admin_role_service"
	"github.com/a20070322/go_fast_admin/app/service/admin_user_service"
	"github.com/a20070322/go_fast_admin/ent"
	"github.com/a20070322/go_fast_admin/global"
	"github.com/go-redis/redis/v8"
	"time"
)

func Init(ctx context.Context) *Cache {
	art := &Cache{}
	art.ctx = ctx
	art.AdminUserPrefix = "admin_user_"
	art.AdminRolePrefix = "admin_role_"
	art.Rdb = global.Rdb
	return art
}

type Cache struct {
	ctx             context.Context
	AdminUserPrefix string
	AdminRolePrefix string
	Rdb             *redis.Client
}

//判断角色是否被缓存
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
	rep := m.Rdb.Set(m.ctx, fmt.Sprintf("%s%s", m.AdminUserPrefix, u.ID), string(str), time.Second*time.Duration(global.AppSetting.Jwt.TokenExpireDuration))
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