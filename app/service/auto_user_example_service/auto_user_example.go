package auto_user_example_service

import (
	"context"
	"errors"
	"github.com/a20070322/go_fast_admin/ent"
	"github.com/a20070322/go_fast_admin/ent/admindict"
	"github.com/a20070322/go_fast_admin/ent/autouserexample"
	"github.com/a20070322/go_fast_admin/ent/predicate"
	"github.com/a20070322/go_fast_admin/global"
	"github.com/a20070322/go_fast_admin/utils/ent_utils"
	"time"
)

//初始化 AutoUserExample service
func Init(ctx context.Context) *AutoUserExample {
	art := &AutoUserExample{}
	art.db = global.Db.AutoUserExample
	art.ctx = ctx
	return art
}

type AutoUserExample struct {
	db  *ent.AutoUserExampleClient
	ctx context.Context
}

// AutoUserExample列表
func (m *AutoUserExample) List(form *FormList) (rep RepList, err error) {
	// 设置分页默认值
	ent_utils.PipePagerFn(form)
	// 查询条件数组
	var whereArr []predicate.AutoUserExample
	// todo 字段搜索条件匹配及过滤匹配
        
        //文本域测试过滤
        if form.TestText != "" {
              whereArr = append(whereArr, autouserexample.TestTextEQ(form.TestText))
            }
        
        //布尔值测试过滤
        if form.TestBool != "" {
            switch form.TestBool {
              case "0":
                  whereArr = append(whereArr, autouserexample.TestBoolEQ(false))
                  break
              case "1":
                  whereArr = append(whereArr, autouserexample.TestBoolEQ(true))
                  break
              }
         }
        
        //数字类型测试过滤
        if form.TestInt > 0 {
              whereArr = append(whereArr, autouserexample.TestIntEQ(form.TestInt))
            }
    // 软删除过滤
    whereArr = append(whereArr, autouserexample.DeletedAtIsNil())
	// 查询
	db := m.db.Query().Where(whereArr...)
	// 获取总条数
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


// AutoUserExample创建
func (m *AutoUserExample) Create(form *FormCreate) (rep RepCreate, err error) {
	db := m.db.
		Create().
	// todo 插入字段过滤
        SetTestText(form.TestText).
        SetTestBool(form.TestBool).
        SetTestInt(form.TestInt)
	u, err := db.Save(m.ctx)
	if err != nil {
		return rep, err
	}
	rep.AutoUserExample = u
	return rep, nil
}

//更新
func (m *AutoUserExample) Update(id int, form *FormUpdate) (rep RepUpdate, err error) {
	fup, err := m.FindById(id)
	if err != nil {
		return rep, errors.New("AutoUserExample is not found")
	}
	db := fup.
		Update().
    // todo 更新字段过滤
        SetTestText(form.TestText).
        SetTestBool(form.TestBool).
        SetTestInt(form.TestInt)
	u, err := db.Save(m.ctx)
	if err != nil {
		return rep, err
	}
	rep.AutoUserExample = u
	return rep, nil
}

//删除
func (m *AutoUserExample) Delete(id int) (err error) {
	fup, err := m.FindById(id)
	if err != nil {
		return errors.New("AutoUserExample is not find")
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
func (m *AdminDict) FindById(id int) (rep *ent.AdminDict, err error) {
	rep, err = m.db.Query().Where(admindict.IDEQ(id), admindict.DeletedAtIsNil()).First(m.ctx)
	if err != nil {
		return rep, errors.New("user is not find")
	}
	return rep, err
}
