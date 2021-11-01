package fast_admin_curd

import (
	"fmt"
	"github.com/a20070322/go_fast_admin/pkg/fast_admin_curd/lib"
	"github.com/a20070322/go_fast_admin/pkg/fast_admin_curd/types"
	"path"
)

func Generate(config *types.ServerConfig) error {
	// 此处并未前端配置，方便后续更改
	config.TmpData.WorkPath=path.Join(config.ProjectConfig.WorkPath,"/pkg/fast_admin_curd")
	w := &lib.Ware{}
	w.Use(lib.HandleModelName)
	//初始化路径
	_, err := w.Run(config)
	if err != nil {
		return err
	}
	fmt.Println("code generation success")
	return nil
}