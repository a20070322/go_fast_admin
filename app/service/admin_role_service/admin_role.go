package admin_role_service

import (
	"context"
	"errors"
	"github.com/a20070322/go_fast_admin/ent"
	"github.com/a20070322/go_fast_admin/ent/adminrole"
	"github.com/a20070322/go_fast_admin/ent/predicate"
	"github.com/a20070322/go_fast_admin/global"
	"github.com/a20070322/go_fast_admin/utils/ent_utils"
	"time"
)

//初始化 AdminRole service
func Init(ctx context.Context) *AdminRole {
	art := &AdminRole{}
	art.db = global.Db.AdminRole
	art.ctx = ctx
	return art
}

type AdminRole struct {
	db  *ent.AdminRoleClient
	ctx context.Context
}

//列表
func (m *AdminRole) List(form *FormList) (rep RepList, err error) {
	// 设置分页默认值
	ent_utils.PipePagerFn(form)
	//查询条件数组
	var whereArr []predicate.AdminRole
	whereArr = append(whereArr, adminrole.DeletedAtIsNil())
	if form.Name != "" {
		whereArr = append(whereArr, adminrole.NameContains(form.Name))
	}
	if form.IsEnable != "" {
		switch form.IsEnable {
		case "0":
			whereArr = append(whereArr, adminrole.IsEnable(false))
			break
		case "1":
			whereArr = append(whereArr, adminrole.IsEnable(true))
			break
		}
	}
	//查询
	db := m.db.Query().Where(whereArr...)
	//获取总条数
	total, err := db.Count(m.ctx)
	if err != nil {
		return rep, err
	}
	rep.Total = total
	// 设置自动分页
	if form.IsAll != "1" {
		ent_utils.PipeLimitFn(db, form)
	}
	rep.Data, err = db.All(m.ctx)
	if err != nil {
		return rep, err
	}
	return rep, nil
}

//创建
func (m *AdminRole) Create(form *FormCreate) (rep RepCreate, err error) {
	db := m.db.
		Create().
		SetName(form.Name).
		SetIsEnable(form.IsEnable)

	u, err := db.Save(m.ctx)
	if err != nil {
		return rep, err
	}
	rep.AdminRole = u
	return rep, nil
}

//更新
func (m *AdminRole) Update(id int, form *FormUpdate) (rep RepUpdate, err error) {
	fup, err := m.FindById(id)
	if err != nil {
		return rep, errors.New("user is not find")
	}
	db := fup.
		Update().
		SetName(form.Name).
		SetIsEnable(form.IsEnable)
	u, err := db.Save(m.ctx)
	if err != nil {
		return rep, err
	}
	rep.AdminRole = u
	return rep, nil
}

//删除
func (m *AdminRole) Delete(id int) (err error) {
	fup, err := m.FindById(id)
	if err != nil {
		return errors.New("AdminRole is not find")
	}
	db := fup.
		Update().
		SetDeletedAt(time.Now())
	_, err = db.Save(m.ctx)
	if err != nil {
		return err
	}
	return nil
}

//查找
func (m *AdminRole) FindById(id int) (rep *ent.AdminRole, err error) {
	rep, err = m.db.Query().Where(adminrole.IDEQ(id), adminrole.DeletedAtIsNil()).First(m.ctx)
	if err != nil {
		return rep, errors.New("user is not find")
	}
	return rep, err
}

func (m *AdminRole) FindByIdWithMenu(id int) (rep *ent.AdminRole, err error) {
	rep, err = m.db.Query().Where(adminrole.IDEQ(id), adminrole.DeletedAtIsNil()).WithMenu().First(m.ctx)
	if err != nil {
		return rep, errors.New("user is not find")
	}
	return rep, err
}

//获取菜单树状权限
func (m *AdminRole) FindMenus(id int) (rep []*ent.AdminMenus, err error) {
	rep, err = m.db.Query().Where(adminrole.IDEQ(id), adminrole.DeletedAtIsNil()).QueryMenu().All(m.ctx)
	if err != nil {
		return rep, errors.New("user is not find")
	}
	return rep, err
}

//设置菜单树状权限
func (m *AdminRole) SetMenus(id int, form *FormSetMenus) (err error) {
	_, err = m.db.Update().Where(adminrole.IDEQ(id), adminrole.DeletedAtIsNil()).ClearMenu().AddMenuIDs(form.Menus...).Save(m.ctx)
	if err != nil {
		return errors.New("user is not find")
	}
	return err
}
