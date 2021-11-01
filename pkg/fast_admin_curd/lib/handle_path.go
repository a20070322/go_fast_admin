package lib

import (
	"errors"
	"fmt"
	"github.com/a20070322/go_fast_admin/pkg/fast_admin_curd/types"
	"github.com/a20070322/go_fast_admin/pkg/fast_admin_curd/utils"
	"path"
)

func HandlePath(config *types.ServerConfig) (*types.ServerConfig, error) {
	if config.ModelName == "" {
		return nil, errors.New("ModelName is empty")
	}

	// controller绝对路径拼接
	config.TmpData.ControllerPath = path.Join(config.ProjectConfig.WorkPath, config.ProjectConfig.ControllerPath)
	config.TmpData.ControllerName = config.TmpData.ModelNameCase + ".go"
	controllerPath := path.Join(config.TmpData.ControllerPath, config.TmpData.ControllerName)
	if !config.GenerateConfig.IsForce && utils.Exists(controllerPath, "file") {
		return config, errors.New(fmt.Sprintf("already exists : %s", controllerPath))
	}

	// service绝对路径拼接
	config.TmpData.ServicePath = path.Join(config.ProjectConfig.WorkPath, config.ProjectConfig.ServicePath,config.TmpData.ModelNameCase+"_service")
	config.TmpData.ServiceName = config.TmpData.ModelNameCase + ".go"
	servicePath := path.Join(config.TmpData.ServicePath, config.TmpData.ServiceName)
	if !config.GenerateConfig.IsForce && utils.Exists(servicePath, "file") {
		return config, errors.New(fmt.Sprintf("already exists : %s", servicePath))
	}

	// service_type绝对路径拼接
	config.TmpData.ServiceTypePath = path.Join(config.ProjectConfig.WorkPath, config.ProjectConfig.ServicePath)
	config.TmpData.ServiceTypeName = "types.go"
	serviceTypePath := path.Join(config.TmpData.ServiceTypePath, config.TmpData.ServiceTypeName)
	if !config.GenerateConfig.IsForce && utils.Exists(serviceTypePath, "file") {
		return config, errors.New(fmt.Sprintf("already exists : %s", serviceTypePath))
	}

	// schema绝对路径拼接
	config.TmpData.SchemaPath = path.Join(config.ProjectConfig.WorkPath, config.ProjectConfig.SchemaPath)
	config.TmpData.SchemaName = config.TmpData.ModelNameLower + ".go"
	schemaPath := path.Join(config.TmpData.SchemaPath, config.TmpData.SchemaName)
	if !config.GenerateConfig.IsForce && utils.Exists(schemaPath, "file") {
		return config, errors.New(fmt.Sprintf("already exists : %s", schemaPath))
	}

	//router文件路径
	config.TmpData.RouterPath = path.Join(config.ProjectConfig.WorkPath, config.ProjectConfig.RouterPath)
	if !utils.Exists(config.TmpData.RouterPath, "file") {
		return config, errors.New(fmt.Sprintf("not found : %s", config.TmpData.RouterPath))
	}

	return config, nil
}
