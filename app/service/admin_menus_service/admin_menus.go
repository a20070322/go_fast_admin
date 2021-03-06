package admin_menus_service

import (
	"context"
	"errors"
	"fmt"
	"github.com/a20070322/go_fast_admin/ent"
	"github.com/a20070322/go_fast_admin/ent/adminmenus"
	"github.com/a20070322/go_fast_admin/ent/predicate"
	"github.com/a20070322/go_fast_admin/global"
	"github.com/a20070322/go_fast_admin/types"
	"github.com/a20070322/go_fast_admin/utils/ent_utils"
	"time"
)

//初始化 AdminMenus service
func Init(ctx context.Context) *AdminMenus {
	art := &AdminMenus{}
	art.db = global.Db.AdminMenus
	art.ctx = ctx
	return art
}

type AdminMenus struct {
	db  *ent.AdminMenusClient
	ctx context.Context
}

//列表
func (m *AdminMenus) List(form *FormList) (rep RepList, err error) {
	// 设置分页默认值
	ent_utils.PipePagerFn(form)
	//查询条件数组
	var whereArr []predicate.AdminMenus
	whereArr = append(whereArr, adminmenus.DeletedAtIsNil())
	//查询
	db := m.db.Query().Where(whereArr...)
	//获取总条数
	total, err := db.Count(m.ctx)
	if err != nil {
		return rep, err
	}
	rep.Total = total
	// 设置自动分页
	ent_utils.PipeLimitFn(db, form)
	rep.Data, err = db.All(m.ctx)
	if err != nil {
		return rep, err
	}
	return rep, nil
}

//列表
func (m *AdminMenus) TreeList(form *FormList) (rep []*types.MenusTree, err error) {
	//查询条件数组
	var whereArr []predicate.AdminMenus
	whereArr = append(whereArr, adminmenus.DeletedAtIsNil())
	//查询
	list, err := m.db.Query().Where(whereArr...).Order(ent.Desc(adminmenus.FieldSort)).All(m.ctx)
	if err != nil {
		return rep, err
	}
	fmt.Println(len(list))
	return MenuToTree(list, 0, 0), nil
}

//创建
func (m *AdminMenus) Create(form *FormCreate) (rep RepCreate, err error) {
	db := m.db.
		Create().
		SetName(form.Name).
		SetPath(form.Path).
		SetRouterPath(form.RouterPath).
		SetIcon(form.Icon).
		SetType(form.Type).
		SetPathAction(form.PathAction).
		SetPowerStr(form.PowerStr).
		SetSort(form.Sort).
		SetFid(form.Fid).
		SetIsExternalLink(form.IsExternalLink).
		SetIsShow(form.IsShow).
		SetIsEnable(form.IsEnable)

	u, err := db.Save(m.ctx)
	if err != nil {
		return rep, err
	}
	rep.AdminMenus = u
	return rep, nil
}

//更新
func (m *AdminMenus) Update(id int, form *FormUpdate) (rep RepUpdate, err error) {
	fup, err := m.FindById(id)
	if err != nil {
		return rep, errors.New("user is not find")
	}
	db := fup.Update().
		SetName(form.Name).
		SetPath(form.Path).
		SetRouterPath(form.RouterPath).
		SetPathAction(form.PathAction).
		SetIcon(form.Icon).
		SetType(form.Type).
		SetPowerStr(form.PowerStr).
		SetSort(form.Sort).
		SetFid(form.Fid).
		SetIsExternalLink(form.IsExternalLink).
		SetIsShow(form.IsShow).
		SetIsEnable(form.IsEnable)

	u, err := db.Save(m.ctx)
	if err != nil {
		return rep, err
	}
	rep.AdminMenus = u
	return rep, nil
}

//删除
func (m *AdminMenus) Delete(id int) (err error) {
	fup, err := m.FindById(id)
	if err != nil {
		return errors.New("AdminMenus is not find")
	}
	count, err := m.db.Query().Where(adminmenus.DeletedAtIsNil(), adminmenus.FidEQ(fup.ID)).Count(m.ctx)
	if err != nil {
		return errors.New("AdminMenus is err")
	}
	if count > 0 {
		return errors.New("存在子菜单，无法删除")
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
func (m *AdminMenus) FindById(id int) (rep *ent.AdminMenus, err error) {
	rep, err = m.db.Query().Where(adminmenus.IDEQ(id), adminmenus.DeletedAtIsNil()).First(m.ctx)
	if err != nil {
		return rep, errors.New("user is not find")
	}
	return rep, err
}

//查找
func (m *AdminMenus) GetWithRole(id int) (rep []*ent.AdminRole, err error) {
	rep, err = m.db.Query().Where(adminmenus.IDEQ(id)).QueryRole().WithMenu(func(query *ent.AdminMenusQuery) {
		query.Where(adminmenus.DeletedAtIsNil())
	}).All(m.ctx)
	if err != nil {
		return rep, errors.New("user is not find")
	}
	return rep, err
}

func MenuToTree(menus []*ent.AdminMenus, fid int, t int) []*types.MenusTree {
	var treeList []*types.MenusTree
	for _, v := range menus {
		if v.Fid == fid && (v.IsShow && v.IsEnable || t != 1) {
			treeList = append(treeList, &types.MenusTree{
				AdminMenus: v,
				Children:   MenuToTree(menus, v.ID, t),
			})
		}
	}
	return treeList
}
