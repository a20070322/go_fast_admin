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

// 生成Schema模板
func GenerateSchema(config *types.ServerConfig) (*types.ServerConfig, error) {
	tpl := template.New("schema.tmpl")
	tpl.Funcs(template.FuncMap{

	})
	// 获取模板位置
	tpl, err := tpl.ParseFiles(path.Join(config.TmpData.WorkPath,"tmpl/schema.tmpl"))
	if err != nil {
		return config, errors.New(fmt.Sprintf("parse template failed, err:%v", err))
	}
	// 判断文件夹是否存在
	err = utils.CheckOrCreateDir(config.TmpData.SchemaPath)
	if err != nil {
		return config, err
	}
	f, err := os.OpenFile(path.Join(config.TmpData.SchemaPath,config.TmpData.SchemaName), os.O_WRONLY|os.O_CREATE, 0644)
	defer f.Close()

	arg := map[string]interface{}{
		"ModelName":config.ModelName,
		"Mixins":config.TableConfig.Mixin,
		"Fields":config.TableConfig.Fields,
	}
	err = tpl.Execute(f, arg)
	if err != nil {
		return config, err
	}
	return config, nil
}