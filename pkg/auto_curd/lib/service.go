package lib

import (
	"fmt"
	"github.com/a20070322/go_fast_admin/pkg/auto_curd/config_type"
	"github.com/a20070322/go_fast_admin/pkg/auto_curd/utils"
	"github.com/pkg/errors"
	"os"
	"path"
	"text/template"
)

func GeneratorService(config *config_type.Config) (*config_type.Config, error) {
	tpl := template.New("service.tmpl")
	tpl.Funcs(template.FuncMap{
		"IsHasUUidFn":     utils.IsHasUUidFn,
		"HumpToLowercase": utils.HumpToLowercase,
		"Case2Camel":      utils.Case2Camel,
		"Camel2Case":      utils.Camel2Case,
		"IdIsUUIDFn":      utils.IdIsUUIDFn,
		"CheckEqLine":     utils.CheckEqLine,
	})
	tpl, err := tpl.ParseFiles(path.Join(config.WorkPath, "tmpl/service.tmpl"))
	if err != nil {
		return config, errors.New(fmt.Sprintf("parse template failed, err:%v", err))
	}
	//var buf bytes.Buffer
	//err = tpl.Execute(&buf, config)
	//if err != nil {
	//	return err
	//}
	//fmt.Println(buf.String())
	//return err
	err = utils.CheckOrCreateDir(config.Runtime.ServicePath)
	if err != nil {
		return config, err
	}
	f, err := os.OpenFile(path.Join(config.Runtime.ServicePath, config.Runtime.ServiceName), os.O_WRONLY|os.O_CREATE, 0644)
	defer f.Close()
	err = tpl.Execute(f, config)
	if err != nil {
		return config, err
	}
	return config, nil
}
