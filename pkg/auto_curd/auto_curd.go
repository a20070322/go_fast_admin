package auto_curd

import (
	"fmt"
	"github.com/a20070322/go_fast_admin/pkg/auto_curd/config_type"
	"github.com/a20070322/go_fast_admin/pkg/auto_curd/lib"
)

func Generate(config *config_type.Config) error {
	config.WorkPath = "./pkg/auto_curd"
	w := &lib.Ware{}
	//初始化路径
	w.Use(lib.InitPath)
	//生成schema
	w.Use(lib.GeneratorSchema)
	//生成service_type
	w.Use(lib.GeneratorServiceTypes)
	//生成service
	w.Use(lib.GeneratorService)
	//生成controller
	w.Use(lib.GeneratorController)
	//注入router
	w.Use(lib.GeneratorRouter)
	_, err := w.Run(config)
	if err != nil {
		return err
	}
	fmt.Println("code generation success")
	return nil
}