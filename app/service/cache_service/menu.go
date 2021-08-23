package cache_service

import (
	"github.com/a20070322/go_fast_admin/app/service/admin_menus_service"
	"github.com/a20070322/go_fast_admin/ent"
	"github.com/a20070322/go_fast_admin/global"
	"github.com/a20070322/go_fast_admin/types"
	"sort"
)

//获取用户菜单
func (m *Cache) GetUserMenu(roleIds []int) (types.RepGetUserMenu, error) {
	var rep types.RepGetUserMenu
	var menus []*ent.AdminMenus
	for _, v := range roleIds {
		role, err := m.GetAdminRoleCatch(v)
		if err != nil {
			return rep, err
		}
		if role.IsEnable == true {
			for _, menu := range role.Edges.Menu {
				if menu.IsEnable == true {
					b := true
					for _, t := range menus {
						if t.ID == menu.ID {
							b = false
						}
					}
					if b {
						//目录及菜单
						if (menu.Type == 1 || menu.Type == 2) && menu.IsShow {
							menus = append(menus, menu)
						}
						//按钮带权限标识
						if menu.Type == 3 || menu.Type == 2{
							rep.Role = append(rep.Role, menu.PowerStr)
						}
					}

				}
			}
		}
	}
	sort.Sort(types.MenuSlice(menus))
	rep.Menu =  admin_menus_service.MenuToTree(menus, 0, 1)
	return rep, nil
}

//菜单更新角色缓存清理
func (m Cache) MenuUpdateRefreshRole(id int) error {
	roles,err := admin_menus_service.Init(m.ctx).GetWithRole(id)
	if err != nil {
		return err
	}
	for _,role := range  roles {
		if m.CheckAdminRoleCatch(role.ID) {
			err = m.SetAdminRoleCatch(role)
			if err != nil {
				global.Logger.Error(err)
			}
		}
	}
	return nil
}