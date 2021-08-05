package casbin_rules_service

import (
	"context"
	"errors"
	"github.com/a20070322/go_fast_admin/ent"
	"github.com/a20070322/go_fast_admin/ent/adminmenus"
	"github.com/a20070322/go_fast_admin/ent/adminrole"
	"github.com/a20070322/go_fast_admin/ent/adminuser"
	"github.com/a20070322/go_fast_admin/ent/casbinrules"
	"github.com/a20070322/go_fast_admin/ent/predicate"
	"github.com/a20070322/go_fast_admin/global"
	"github.com/a20070322/go_fast_admin/utils/ent_utils"
	"strconv"
)

//初始化 CasbinRules service
func Init(ctx context.Context) *CasbinRules {
	art := &CasbinRules{}
	art.db = global.Db.CasbinRules
	art.ctx = ctx
	return art
}

type CasbinRules struct {
	db  *ent.CasbinRulesClient
	ctx context.Context
}

//列表
func (m *CasbinRules) List(form *FormList) (rep RepList, err error) {
	// 设置分页默认值
	ent_utils.PipePagerFn(form)
	//查询条件数组
	var whereArr []predicate.CasbinRules
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

//创建
func (m *CasbinRules) Create(form *FormCreate) (rep RepCreate, err error) {
	db := m.db.
		Create().
		SetPtype(form.Ptype).
		SetV0(form.V0).
		SetV1(form.V1).
		SetV2(form.V2).
		SetV3(form.V3).
		SetV4(form.V4).
		SetV5(form.V5)

	u, err := db.Save(m.ctx)
	if err != nil {
		return rep, err
	}
	rep.CasbinRules = u
	return rep, nil
}

//更新
func (m *CasbinRules) Update(id int, form *FormUpdate) (rep RepUpdate, err error) {
	fup, err := m.FindById(id)
	if err != nil {
		return rep, errors.New("user is not find")
	}
	db := fup.
		Update().
		SetPtype(form.Ptype).
		SetV0(form.V0).
		SetV1(form.V1).
		SetV2(form.V2).
		SetV3(form.V3).
		SetV4(form.V4).
		SetV5(form.V5)

	u, err := db.Save(m.ctx)
	if err != nil {
		return rep, err
	}
	rep.CasbinRules = u
	return rep, nil
}

//删除
func (m *CasbinRules) Delete(id int) (err error) {
	fup, err := m.FindById(id)
	if err != nil {
		return errors.New("CasbinRules is not find")
	}
	err = m.db.DeleteOne(fup).Exec(m.ctx)
	if err != nil {
		return err
	}
	return nil
}

//查找
func (m *CasbinRules) FindById(id int) (rep *ent.CasbinRules, err error) {
	rep, err = m.db.Query().Where(casbinrules.IDEQ(id)).First(m.ctx)
	if err != nil {
		return rep, errors.New("user is not find")
	}
	return rep, err
}

func (m *CasbinRules) InitCasbinEnt() error {
	roles, err := global.Db.AdminRole.Query().
		Where(adminrole.DeletedAtIsNil()).
		WithMenu(func(query *ent.AdminMenusQuery) {
			query.Where(adminmenus.TypeNEQ(1), adminmenus.DeletedAtIsNil())
		}).WithUser(func(query *ent.AdminUserQuery) {
			query.Where(adminuser.DeletedAtIsNil())
		}).All(m.ctx)
	if err != nil {
		return err
	}
	var bulk []*ent.CasbinRulesCreate
	for _, r := range roles {
		//创建用户与角色关联
		for _,u := range r.Edges.User{
			bulk = append(bulk, m.db.Create().SetPtype("g").SetV0(u.ID.String()).SetV1(strconv.Itoa(r.ID)))
		}
		//创建角色与menu
		for _,menu := range r.Edges.Menu{
			bulk = append(bulk, m.db.Create().SetPtype("p").SetV0(strconv.Itoa(r.ID)).SetV1(menu.Path).SetV2(menu.PathAction))
		}
	}
	_, err = m.db.CreateBulk(bulk...).Save(m.ctx)
	return err
}
