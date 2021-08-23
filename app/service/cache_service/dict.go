package cache_service

import (
	"encoding/json"
	"errors"
	"github.com/a20070322/go_fast_admin/app/service/admin_dict_service"
	"github.com/a20070322/go_fast_admin/global"
)

//设置dict缓存
func (m Cache) SetDictMapCatch(u map[string][]admin_dict_service.DictMap) error {
	if u == nil {
		return errors.New("ent.AdminRole is nil")
	}
	global.Logger.Debug("字典缓存更新")
	str, err := json.Marshal(u)
	if err != nil {
		global.Logger.Error(err)
		return err
	}
	rep := m.Rdb.Set(m.ctx, "dic_map", string(str),0)
	if rep.Err() != nil {
		global.Logger.Error(rep.Err())
		return rep.Err()
	}
	return nil
}
//读取dict缓存
func (m Cache) GetDictMapCatch() (map[string][]admin_dict_service.DictMap, error) {
	rep, err := m.Rdb.Get(m.ctx, "dic_map").Result()
	if err != nil {
		global.Logger.Error(err)
		u, err := admin_dict_service.Init(m.ctx).GetDictMap()
		_ = m.SetDictMapCatch(u)
		if err != nil {
			global.Logger.Error(err)
			return nil, err
		}
		return u, err
	}
	var u = make(map[string][]admin_dict_service.DictMap)
	err = json.Unmarshal([]byte(rep), &u)
	if err != nil {
		global.Logger.Error(err)
		return nil, err
	}
	return u, nil
}
