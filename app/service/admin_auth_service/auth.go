package admin_auth_service

import (
	"context"
	"errors"
	"github.com/a20070322/go_fast_admin/app/service/cache_service"
	"github.com/a20070322/go_fast_admin/ent"
	"github.com/a20070322/go_fast_admin/ent/adminuser"
	"github.com/a20070322/go_fast_admin/global"
	"github.com/a20070322/go_fast_admin/utils/jwt"
	"github.com/a20070322/go_fast_admin/utils/pass"
)

func Init(ctx context.Context) *Auth {
	art := &Auth{}
	art.db = global.Db.AdminUser
	art.ctx = ctx
	return art
}

type Auth struct {
	db  *ent.AdminUserClient
	ctx context.Context
}

//登录
func (m *Auth) Login(form *FormLogin) (*RepGetToken, error) {
	u, err := m.db.Query().Where(adminuser.UsernameEQ(form.Username)).WithRole().First(m.ctx)
	if err != nil {
		return nil, errors.New("用户名或密码错误_001")
	}
	if u.IsEnable == false {
		return nil, errors.New("该用户已被禁用")
	}
	pErr := pass.DecryptPassword(u.Password, form.Password)
	if pErr != nil {
		return nil, errors.New("用户名或密码错误_002")
	}
	return m.GetToken(u)
}

//获取token
func (m *Auth) GetToken(user *ent.AdminUser) (*RepGetToken, error) {
	var rep RepGetToken
	rep.User = user
	j, jErr := jwt.GenToken(&jwt.Claims{
		UserGroup: jwt.UserGroupAdmin,
		UserID:    user.ID.String(),
	})
	if jErr != nil {
		return nil, errors.New("token 生成异常")
	}
	rep.JwtData = j
	//重新缓存用户
	cache_service.Init(m.ctx).SetAdminUserCatch(user)
	return &rep, nil
}
