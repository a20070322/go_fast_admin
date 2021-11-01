package test

import (
	"github.com/a20070322/go_fast_admin/pkg/fast_admin_curd/lib"
	"github.com/a20070322/go_fast_admin/pkg/fast_admin_curd/utils"
	"path"
	"testing"
)

func TestGetConfig(t *testing.T) {
	config,err := utils.GetConfig("/Users/zhaozhongyang/Desktop/goAdmin/pkg/fast_admin_curd/test/config.json")
	if err != nil {
		t.Error(err)
	}
	if config.ModelName != "AutoCurdTest" {
		t.Error("ModelName is not AutoCurdTest")
	}
}

func TestHandleModelName(t *testing.T) {
	config,err := utils.GetConfig("/Users/zhaozhongyang/Desktop/goAdmin/pkg/fast_admin_curd/test/config.json")
	if err != nil {
		t.Error(err)
	}
	w := &lib.Ware{}
	w.Use(lib.HandleModelName)
	//初始化路径
	config, err = w.Run(config)
	if err != nil {
		t.Error(err)
	}
	if "autocurdtest" != config.TmpData.ModelNameLower{
		t.Error("ModelNameLower is bad")
	}
	if "auto_curd_test" != config.TmpData.ModelNameCase{
		t.Error("ModelNameCase is bad")
	}
}

// 测试生成器
func TestGenerate(t *testing.T) {
	config,err := utils.GetConfig("/Users/zhaozhongyang/Desktop/开源项目/goAdmin/pkg/fast_admin_curd/test/config.json")
	if err != nil {
		t.Error(err)
	}
	config.TmpData.WorkPath=path.Join(config.ProjectConfig.WorkPath,"/pkg/fast_admin_curd")
	w := &lib.Ware{}
	w.Use(lib.HandleModelName)
	w.Use(lib.HandlePath)
	w.Use(lib.GenerateSchema)
	w.Use(lib.GenerateService)
	//初始化路径
	config, err = w.Run(config)
	if err != nil {
		t.Error(err)
	}
}
