package admin_dict_key_service

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

//初始化 AdminDictKey service
func Init(ctx context.Context) *AdminDictKey {
	art := &AdminDictKey{}
	art.db = global.Db.AdminDictKey
	art.ctx = ctx
	return art
}

type AdminDictKey struct {
	db  *ent.AdminDictKeyClient
	ctx context.Context
}


//列表
func (m *AdminDictKey) List(form *FormList) (rep RepList, err error) {
	// 设置分页默认值
	ent_utils.PipePagerFn(form)
	//查询条件数组
	var whereArr []predicate.AdminDictKey
	whereArr = append(whereArr, admindictkey.DeletedAtIsNil())
	global.Logger.Debug(form.Fid)
	if form.Fid > 0 {
		whereArr = append(whereArr,admindictkey.HasPWith(admindict.IDEQ(form.Fid)))
	}
	//查询
	db := m.db.Query().Where(whereArr...).WithP().Order(ent.Desc(admindictkey.FieldSort))
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
func (m *AdminDictKey) Create(form *FormCreate) (rep RepCreate, err error) {
	db := m.db.
		Create().
		SetDictLabel(form.DictLabel).
        SetDictCode(form.DictCode).
        SetSort(form.Sort).
        SetRemarks(form.Remarks).
        SetIsEnable(form.IsEnable).
		SetPID(form.Fid)
        
	u, err := db.Save(m.ctx)
	if err != nil {
		return rep, err
	}
	rep.AdminDictKey = u
	return rep, nil
}

//更新
func (m *AdminDictKey) Update(id int , form *FormUpdate) (rep RepUpdate, err error) {
	fup, err := m.FindById(id)
	if err != nil {
		return rep, errors.New("user is not find")
	}
	db := fup.
		Update().
		SetDictLabel(form.DictLabel).
        SetDictCode(form.DictCode).
        SetSort(form.Sort).
        SetRemarks(form.Remarks).
        SetIsEnable(form.IsEnable).
		SetPID(form.Fid)
        
	u, err := db.Save(m.ctx)
	if err != nil {
		return rep, err
	}
	rep.AdminDictKey = u
	return rep, nil
}

//删除
func (m *AdminDictKey) Delete(id int ) (err error) {
    fup, err := m.FindById(id)
	if err != nil {
		return errors.New("AdminDictKey is not find")
	}
	err =  m.db.DeleteOneID(fup.ID).Exec(m.ctx)
	if err != nil {
		return err
	}
	return nil
}

//查找
func (m *AdminDictKey) FindById(id int ) (rep *ent.AdminDictKey, err error) {
	rep, err = m.db.Query().Where(admindictkey.IDEQ(id),admindictkey.DeletedAtIsNil()).First(m.ctx)
	if err != nil {
		return rep, errors.New("user is not find")
	}
	return rep, err
}