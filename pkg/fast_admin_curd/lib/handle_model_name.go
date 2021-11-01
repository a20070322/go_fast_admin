package lib

import (
	"errors"
	"github.com/a20070322/go_fast_admin/pkg/fast_admin_curd/types"
	"github.com/a20070322/go_fast_admin/pkg/fast_admin_curd/utils"
	"strings"
)

func HandleModelName(config *types.ServerConfig) (*types.ServerConfig, error) {
	if config.ModelName == "" {
		return nil, errors.New("ModelName is empty")
	}
	config.TmpData.ModelNameCase = utils.Camel2Case(config.ModelName)
	config.TmpData.ModelNameLower = strings.ToLower(config.ModelName)
	return config, nil
}

//func InitPath(config *config_type.Config) (*config_type.Config, error) {
