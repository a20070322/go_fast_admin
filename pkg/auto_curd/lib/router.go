package lib

import (
	"bytes"
	"fmt"
	"github.com/a20070322/go_fast_admin/pkg/auto_curd/config_type"
	"github.com/a20070322/go_fast_admin/pkg/auto_curd/utils"
	"os"
	"path"
	"text/template"
)

type InjectionRouter struct {
	RouteTmpl string `json:"route_tmpl"`
}
func GeneratorRouter(config *config_type.Config) (*config_type.Config, error) {
	tpl := template.New("router.tmpl")
	tpl.Funcs(template.FuncMap{
		"IsHasUUidFn":         utils.IsHasUUidFn,
		"HumpToLowercase":     utils.HumpToLowercase,
		"Case2Camel":          utils.Case2Camel,
		"Camel2Case":          utils.Camel2Case,
		"GetControllerModule": utils.GetControllerModule,
	})

	tpl, err := tpl.ParseFiles(path.Join(config.WorkPath,"tmpl/router.tmpl"))
	if err != nil {
		fmt.Printf("parse template failed, err:%v\n", err)
	}
	var buf bytes.Buffer
	err = tpl.Execute(&buf, config)
	if err != nil {
		return config,err
	}
	//router 自动注入
	rTpl := template.New(utils.GetControllerModule(config.ProjectConfig.RouterPath))
	rTpl, err = rTpl.ParseFiles(config.Runtime.RouterPath)
	f, err := os.OpenFile(config.Runtime.RouterPath, os.O_WRONLY|os.O_CREATE, 0644)
	defer f.Close()
	err = rTpl.Execute(f,&InjectionRouter{
		RouteTmpl:buf.String(),
	})
	if err != nil {
		fmt.Println(err)
		return config, err
	}
	return config,err
}
