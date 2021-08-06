package check_role_service

import (
	"context"
	"github.com/a20070322/go_fast_admin/app/service/cache_service"
	"github.com/a20070322/go_fast_admin/global"
	"regexp"
	"strings"
)

func RegexMatch(key1 string, key2 string) bool {
	res, err := regexp.MatchString(key2, key1)
	if err != nil {
		panic(err)
	}
	return res
}

// /abc/:id  /abc/*  /abc/abc  /abc/:id/abc
func KeyMatch(key1 string, key2 string) bool {
	key2 = strings.Replace(key2, "/*", "/.*", -1)

	re := regexp.MustCompile(`:[^/]+`)
	key2 = re.ReplaceAllString(key2, "$1[^/]+$2")

	return RegexMatch(key1, "^"+key2+"$")
}

type FormAdminCheckRole struct {
	Uid    string `json:"uid"`
	Path   string `json:"path"`
	Method string `json:"method"`
}

func AdminCheckRole(ctx context.Context, form *FormAdminCheckRole) (bool, error) {
	catch := cache_service.Init(ctx)
	u, err := catch.GetAdminUserCatch(form.Uid)
	if err != nil {
		global.Logger.Error(err)
		return false, err
	}
	for _, v := range u.Edges.Role {
		m, err := catch.GetAdminRoleCatch(v.ID)
		if err != nil {
			global.Logger.Error(err)
			return false, err
		}
		if m.IsEnable != true {
			break
		}
		for _, r := range m.Edges.Menu {
			if r.PathAction == form.Method && KeyMatch(form.Path, r.Path) && r.IsEnable {
				return true, nil
			}
		}
	}
	return false, nil
}
