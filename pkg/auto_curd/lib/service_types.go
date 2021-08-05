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

func GeneratorServiceTypes(config *config_type.Config) (*config_type.Config, error) {
	tpl := template.New("service_types.tmpl")
	tpl.Funcs(template.FuncMap{
		"IsHasUUidFn":     utils.IsHasUUidFn,
		"HumpToLowercase": utils.HumpToLowercase,
		"Case2Camel": utils.Case2Camel,
		"Camel2Case":      utils.Camel2Case,
		"IdIsUUIDFn": func(fields []config_type.FieldType) (b bool) {
			b = false
			for _, v := range fields {
				if v.Name == "id" && v.Type == "UUID" {
					b = true
					break
				}
			}
			return b
		},
	})
	tpl, err := tpl.ParseFiles(path.Join(config.WorkPath,"tmpl/service_types.tmpl"))
	if err != nil {
		return config, errors.New(fmt.Sprintf("parse template failed, err:%v", err))
	}

	//var buf bytes.Buffer
	//err = tpl.Execute(&buf, config)
	//if err != nil {
	//	return config, err
	//}
	//fmt.Println(buf.String())
	//return config,err

	err = utils.CheckOrCreateDir(config.Runtime.ServiceTypePath)
	if err != nil {
		return config, err
	}
	f, err := os.OpenFile(path.Join(config.Runtime.ServiceTypePath,config.Runtime.ServiceTypeName), os.O_WRONLY|os.O_CREATE, 0644)
	defer f.Close()
	err = tpl.Execute(f, config)
	if err != nil {
		return config, err
	}
	return config, nil
}
