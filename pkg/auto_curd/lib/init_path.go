package lib

import (
	"fmt"
	"github.com/a20070322/go_fast_admin/pkg/auto_curd/config_type"
	"github.com/a20070322/go_fast_admin/pkg/auto_curd/utils"
	"github.com/pkg/errors"
	"os"
	"path"
)

func InitPath(config *config_type.Config) (*config_type.Config, error) {
	var err error
	//项目绝对路径设置
	config.Runtime.ProjectPath, err = os.Getwd()
	if err != nil {
		return config, err
	}
	//controller绝对路径拼接
	config.Runtime.ControllerPath = path.Join(config.Runtime.ProjectPath, config.ProjectConfig.ControllerPath)
	config.Runtime.ControllerName = utils.Camel2Case(config.Table.Name) + ".go"

	if !config.ProjectConfig.IsForce && utils.Exists(path.Join(config.Runtime.ControllerPath, config.Runtime.ControllerName), "file") {
		return config, errors.New(fmt.Sprintf("already exists : %s", config.Runtime.ControllerPath))
	}
	//service文件夹路径
	serviceDir := path.Join(config.Runtime.ProjectPath, config.ProjectConfig.ServicePath, utils.Camel2Case(config.Table.Name)+"_service")

	//service
	config.Runtime.ServicePath = path.Join(serviceDir)
	config.Runtime.ServiceName = utils.Camel2Case(config.Table.Name) + ".go"
	if !config.ProjectConfig.IsForce &&utils.Exists(path.Join(config.Runtime.ServicePath, config.Runtime.ServiceName), "file") {
		return config, errors.New(fmt.Sprintf("already exists : %s", config.Runtime.ServicePath))
	}
	//service_type
	config.Runtime.ServiceTypePath = path.Join(serviceDir)
	config.Runtime.ServiceTypeName = "types.go"
	if !config.ProjectConfig.IsForce &&utils.Exists(path.Join(config.Runtime.ServiceTypePath, config.Runtime.ServiceTypeName), "file") {
		return config, errors.New(fmt.Sprintf("already exists : %s", config.Runtime.ServiceTypePath))
	}

	//schema路径
	config.Runtime.SchemaPath = path.Join(config.Runtime.ProjectPath, config.ProjectConfig.SchemaPath)
	config.Runtime.SchemaName = utils.HumpToLowercase(config.Table.Name) + ".go"
	if !config.ProjectConfig.IsForce &&utils.Exists(path.Join(config.Runtime.SchemaPath, config.Runtime.SchemaName), "file") {
		return config, errors.New(fmt.Sprintf("already exists : %s", config.Runtime.SchemaPath))
	}

	//router文件路径
	config.Runtime.RouterPath = path.Join(config.Runtime.ProjectPath, config.ProjectConfig.RouterPath)
	if !utils.Exists(config.Runtime.RouterPath, "file") {
		return config, errors.New(fmt.Sprintf("not found : %s", config.Runtime.RouterPath))
	}

	return config, nil
}
