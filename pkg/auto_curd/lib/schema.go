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

func GeneratorSchema(config *config_type.Config) (*config_type.Config, error) {
	tpl := template.New("schema.tmpl")
	tpl.Funcs(template.FuncMap{
		"IsHasUUidFn": utils.IsHasUUidFn,
	})
	tpl, err := tpl.ParseFiles(path.Join(config.WorkPath,"tmpl/schema.tmpl"))
	if err != nil {
		return config, errors.New(fmt.Sprintf("parse template failed, err:%v", err))
	}
	//var buf bytes.Buffer
	//err = tpl.Execute(&buf,config.Table)
	//if err != nil {
	//	return config,err
	//}
	//fmt.Println(buf.String())
	err = utils.CheckOrCreateDir(config.Runtime.SchemaPath)
	if err != nil {
		return config, err
	}
	f, err := os.OpenFile(path.Join(config.Runtime.SchemaPath,config.Runtime.SchemaName), os.O_WRONLY|os.O_CREATE, 0644)
	defer f.Close()
	err = tpl.Execute(f, config.Table)
	if err != nil {
		return config, err
	}
	return config, nil
}
