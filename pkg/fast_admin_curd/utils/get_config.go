package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/a20070322/go_fast_admin/pkg/fast_admin_curd/types"
	"github.com/cjrd/allocate"
	"os"
)

func GetConfig(path string) (*types.ServerConfig, error) {
	var config types.ServerConfig
	filePtr, err := os.Open(path)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Open file failed [Err:%s]", err.Error()))
	}
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}
	//初始化
	allocate.Zero(&config)
	defer filePtr.Close()
	return &config,nil
}