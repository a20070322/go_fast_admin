package bootstrap

import (
	"fmt"
	"github.com/a20070322/go_fast_admin/global"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/ent-adapter"
	"github.com/labstack/gommon/color"
)

func CasbinInit() {
	c := global.AppSetting.Database
	a, err := entadapter.NewAdapter("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.User, c.Password, c.Host, c.Port, c.Dbname))
	if err != nil {
		panic(err)
	}
	e, err := casbin.NewEnforcer("./config/rbac.conf", a)
	if err != nil {
		panic(err)
	}
	global.Rbac = e
	fmt.Println("go_fast_admin: " + color.Green("casbin初始化成功"))
}
