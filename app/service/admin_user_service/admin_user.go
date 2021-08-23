package admin_user_service

import (
	"context"
	"errors"
	"github.com/a20070322/go_fast_admin/ent"
	"github.com/a20070322/go_fast_admin/ent/adminrole"
	"github.com/a20070322/go_fast_admin/ent/adminuser"
	"github.com/a20070322/go_fast_admin/ent/predicate"
	"github.com/a20070322/go_fast_admin/global"
	"github.com/a20070322/go_fast_admin/utils/ent_utils"
	"github.com/a20070322/go_fast_admin/utils/pass"
	"github.com/google/uuid"
	"time"
)

func Init(ctx context.Context) *AdminUser {
	art := &AdminUser{}
	art.db = global.Db.AdminUser
	art.ctx = ctx
	return art
}

type AdminUser struct {
	db  *ent.AdminUserClient
	ctx context.Context
}

//列表
func (m *AdminUser) List(form *FromList) (rep RepList, err error) {
	// 设置分页默认值
	ent_utils.PipePagerFn(form)
	//查询条件数组
	var whereArr []predicate.AdminUser
	whereArr = append(whereArr, adminuser.DeletedAtIsNil())
	if form.Username != "" {
		whereArr = append(whereArr, adminuser.UsernameContains(form.Username))
	}
	if form.Phone != "" {
		whereArr = append(whereArr, adminuser.PhoneContains(form.Phone))
	}
	if form.IsEnable != "" {
		switch form.IsEnable {
		case "0":
			whereArr = append(whereArr, adminuser.IsEnable(false))
			break
		case "1":
			whereArr = append(whereArr, adminuser.IsEnable(true))
			break
		}
	}
	if form.Role != 0 {
		whereArr = append(whereArr,adminuser.HasRoleWith(adminrole.IDEQ(form.Role)))
	}
	//查询
	db := m.db.Query().Where(whereArr...).WithRole(func(query *ent.AdminRoleQuery){
		query.Select(adminrole.FieldID, adminrole.FieldName)
	})
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
func (m *AdminUser) Create(form *FromCommon) (rep RepCommon, err error) {
	form.Password, err = pass.EncryptPassword(form.Password)
	if err != nil {
		return rep, err
	}
	db := m.db.
		Create().
		SetUsername(form.Username).
		SetPassword(form.Password).
		SetAvatar(form.Avatar).
		SetPhone(form.Phone).
		SetIsEnable(form.IsEnable).
		AddRoleIDs(form.Roles...)
	u, err := db.Save(m.ctx)
	if err != nil {
		return rep, err
	}
	rep.AdminUser = u
	return rep, nil
}

//更新
func (m *AdminUser) Update(id string, form *FromUpdate) (rep RepCommon, err error) {
	fup, err := m.FindById(id)
	if err != nil {
		return rep, errors.New("user is not find")
	}
	db := fup.
		Update().
		SetUsername(form.Username).
		//SetPassword(form.Password).
		SetAvatar(form.Avatar).
		SetPhone(form.Phone).
		SetIsEnable(form.IsEnable).
		ClearRole().
		AddRoleIDs(form.Roles...)
	u, err := db.Save(m.ctx)
	if err != nil {
		return rep, err
	}
	rep.AdminUser = u
	return rep, nil
}

//删除
func (m *AdminUser) Delete(id string) (err error) {
	fup, err := m.FindById(id)
	if err != nil {
		return errors.New("user is not find")
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
func (m *AdminUser) FindById(id string) (rep *ent.AdminUser, err error) {
	rep, err = m.db.Query().Where(adminuser.IDEQ(uuid.MustParse(id)), adminuser.DeletedAtIsNil()).WithRole().First(m.ctx)
	if err != nil {
		return rep, errors.New("user is not find")
	}
	return rep, err
}

//更新密码
func (m *AdminUser) UpdatePass(id string, passStr string) (err error) {
	fup, err := m.db.Query().Where(adminuser.IDEQ(uuid.MustParse(id)), adminuser.DeletedAtIsNil()).First(m.ctx)
	if err != nil {
		return errors.New("user is not find")
	}
	passStr, err = pass.EncryptPassword(passStr)
	if err != nil {
		return err
	}
	db := fup.
		Update().
		SetPassword(passStr)
	_, err = db.Save(m.ctx)
	if err != nil {
		return err
	}
	return nil
}
