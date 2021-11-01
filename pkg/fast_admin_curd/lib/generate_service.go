package lib

import (
	"errors"
	"fmt"
	"github.com/a20070322/go_fast_admin/pkg/fast_admin_curd/types"
	"github.com/a20070322/go_fast_admin/pkg/fast_admin_curd/utils"
	"os"
	"path"
	"text/template"
)
// 生成service模板
func GenerateService(config *types.ServerConfig) (*types.ServerConfig, error) {
	tpl := template.New("service.tmpl")
	tpl.Funcs(template.FuncMap{
		"GetFiledType": func( filedType string ) string {
			if utils.In(filedType,[]string{"Int", "Int8", "Int16","Int32","Int64","Float","Float32"}) {
				return  "number"
			}
			if utils.In(filedType,[]string{"String", "Text","Time","UUID"}) {
				return  "string"
			}
			if utils.In(filedType,[]string{"Bool"}) {
				return  "bool"
			}
			return "none"
		},
		"Case2Camel":utils.Case2Camel,
		"Len": func( fields []types.FieldsItem) int {
			return len(fields) - 1
		},
	})
	// 获取模板位置
	tpl, err := tpl.ParseFiles(path.Join(config.TmpData.WorkPath,"tmpl/service.tmpl"))
	if err != nil {
		return config, errors.New(fmt.Sprintf("parse template failed, err:%v", err))
	}
	// 判断文件夹是否存在
	err = utils.CheckOrCreateDir(config.TmpData.ServicePath)
	if err != nil {
		return config, err
	}
	f, err := os.OpenFile(path.Join(config.TmpData.ServicePath,config.TmpData.ServiceName), os.O_WRONLY|os.O_CREATE, 0644)
	defer f.Close()

	arg := map[string]interface{}{
		"ModelName":config.ModelName,
		"ModelNameCase":config.TmpData.ModelNameCase,
		"ModulePath":config.ProjectConfig.ModulePath,
		"ModelNameLower":config.TmpData.ModelNameLower,
		"Fields":config.TableConfig.Fields,
		"IsSoftDel":config.TableConfig.IsSoftDel,
	}
	err = tpl.Execute(f, arg)
	if err != nil {
		return config, err
	}
	return config, nil
}