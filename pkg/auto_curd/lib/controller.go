package lib

import (
	"errors"
	"fmt"
	"github.com/a20070322/go_fast_admin/pkg/auto_curd/config_type"
	"github.com/a20070322/go_fast_admin/pkg/auto_curd/utils"
	"os"
	"path"
	"text/template"
)

func GeneratorController(config *config_type.Config) (*config_type.Config, error) {
	tpl := template.New("controller.tmpl")
	tpl.Funcs(template.FuncMap{
		"IdIsUUIDFn":utils.IdIsUUIDFn,
		"IsHasUUidFn": utils.IsHasUUidFn,
		"GetControllerModule":utils.GetControllerModule,
		"Camel2Case":utils.Camel2Case,
	})
	tpl, err := tpl.ParseFiles(path.Join(config.WorkPath,"tmpl/controller.tmpl"))
	if err != nil {
		return config, errors.New(fmt.Sprintf("parse template failed, err:%v", err))
	}
	//var buf bytes.Buffer
	//err = tpl.Execute(&buf,config)
	//if err != nil {
	//	return err
	//}
	//fmt.Println(buf.String())
	//f, err := os.OpenFile("./dist/test.go", os.O_WRONLY|os.O_CREATE, 0644)
	//defer f.Close()
	//err = tpl.Execute(f, table)
	//if err != nil {
	//	return err
	//}
	err = utils.CheckOrCreateDir(config.Runtime.ControllerPath)
	if err != nil {
		return config, err
	}
	f, err := os.OpenFile(path.Join(config.Runtime.ControllerPath,config.Runtime.ControllerName), os.O_WRONLY|os.O_CREATE, 0644)
	defer f.Close()
	err = tpl.Execute(f, config)
	if err != nil {
		return config, err
	}
	return config, nil
}

