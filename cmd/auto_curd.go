package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/a20070322/go_fast_admin/pkg/auto_curd"
	"github.com/a20070322/go_fast_admin/pkg/auto_curd/config_type"
	"github.com/cjrd/allocate"
	"github.com/labstack/gommon/color"
	"github.com/urfave/cli/v2"
	"os"
)
func getConfig(path string) (*config_type.Config, error) {
	var config config_type.Config
	filePtr, err := os.Open(path)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Open file failed [Err:%s]", err.Error()))
	}
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&config)
	filePtr.Close()
	if err != nil {
		return nil, err
	}
	//初始化
	allocate.Zero(&config)
	return &config,nil
}

var AutoCurd = &cli.Command{
	Name:  "AutoCurd",
	Usage: "自动生成curd代码，及路由注入",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "conf",
			Usage:   "配置文件",
			Aliases: []string{"c"},
		},
	},
	Action: func(ctx *cli.Context) error {
		conf := ctx.String("conf")
		fmt.Println(fmt.Sprintf("配置文件: %s",color.Yellow(conf)) )
		config,err := getConfig(conf)
		if err != nil {
			return err
		}
		return auto_curd.Generate(config)
	},
}
