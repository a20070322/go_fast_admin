package admin_dict_service

import (
	"context"
	"errors"
	"github.com/a20070322/go_fast_admin/ent"
	"github.com/a20070322/go_fast_admin/ent/admindict"
	"github.com/a20070322/go_fast_admin/ent/admindictkey"
	"github.com/a20070322/go_fast_admin/ent/predicate"
	"github.com/a20070322/go_fast_admin/global"
	"github.com/a20070322/go_fast_admin/utils/ent_utils"
)

//初始化 AdminDict service
func Init(ctx context.Context) *AdminDict {
	art := &AdminDict{}
	art.db = global.Db.AdminDict
	art.ctx = ctx
	return art
}

type AdminDict struct {
	db  *ent.AdminDictClient
	ctx context.Context
}

//列表
func (m *AdminDict) List(form *FormList) (rep RepList, err error) {
	// 设置分页默认值
	ent_utils.PipePagerFn(form)
	//查询条件数组
	var whereArr []predicate.AdminDict
	if form.IsEnable != "" {
		switch form.IsEnable {
		case "0":
			whereArr = append(whereArr, admindict.IsEnable(false))
			break
		case "1":
			whereArr = append(whereArr, admindict.IsEnable(true))
			break
		}
	}
	if form.DictName != "" {
		whereArr = append(whereArr, admindict.DictNameContains(form.DictName))
	}
	if form.DictType != "" {
		whereArr = append(whereArr, admindict.DictTypeContains(form.DictType))
	}

	whereArr = append(whereArr, admindict.DeletedAtIsNil())
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
func (m *AdminDict) Create(form *FormCreate) (rep RepCreate, err error) {
	db := m.db.
		Create().
		SetDictType(form.DictType).
		SetDictName(form.DictName).
		SetRemarks(form.Remarks).
		SetIsEnable(form.IsEnable)

	u, err := db.Save(m.ctx)
	if err != nil {
		return rep, err
	}
	rep.AdminDict = u
	return rep, nil
}

//更新
func (m *AdminDict) Update(id int, form *FormUpdate) (rep RepUpdate, err error) {
	fup, err := m.FindById(id)
	if err != nil {
		return rep, errors.New("user is not find")
	}
	db := fup.
		Update().
		SetDictType(form.DictType).
		SetDictName(form.DictName).
		SetRemarks(form.Remarks).
		SetIsEnable(form.IsEnable)

	u, err := db.Save(m.ctx)
	if err != nil {
		return rep, err
	}
	rep.AdminDict = u
	return rep, nil
}

//删除
func (m *AdminDict) Delete(id int) (err error) {
	fup, err := m.FindById(id)
	if err != nil {
		return errors.New("AdminDict is not find")
	}
	//db := fup.
	//	Update().
	//	SetDeletedAt(time.Now())
	//_, err = db.Save(m.ctx)
	_, err = global.Db.AdminDictKey.Delete().Where(admindictkey.HasPWith(admindict.ID(fup.ID))).Exec(m.ctx)
	if err != nil {
		return err
	}
	err = m.db.DeleteOneID(fup.ID).Exec(m.ctx)
	if err != nil {
		return err
	}
	return nil
}

//查找
func (m *AdminDict) FindById(id int) (rep *ent.AdminDict, err error) {
	rep, err = m.db.Query().Where(admindict.IDEQ(id), admindict.DeletedAtIsNil()).First(m.ctx)
	if err != nil {
		return rep, errors.New("user is not find")
	}
	return rep, err
}

//生成字典集合
func (m *AdminDict) GetDictMap() (map[string][]DictMap, error) {
	repMap := make(map[string][]DictMap)
	dict, err := m.db.Query().Where(admindict.IsEnable(true), admindict.DeletedAtIsNil()).WithKey(func(query *ent.AdminDictKeyQuery) {
		query.Order(ent.Desc(admindictkey.FieldSort))
	}).All(m.ctx)
	if err != nil {
		return repMap, err
	}
	for _, k := range dict {
		repMap[k.DictType] = make([]DictMap,0)
		for _, v := range k.Edges.Key {
			repMap[k.DictType] = append(repMap[k.DictType], DictMap{
				Label: v.DictLabel,
				Value:  v.DictCode,
			})
		}
	}
	return repMap, err
}
