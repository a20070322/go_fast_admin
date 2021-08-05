package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/a20070322/go_fast_admin/bootstrap"
	"github.com/a20070322/go_fast_admin/ent"
	"github.com/a20070322/go_fast_admin/ent/migrate"
	"github.com/a20070322/go_fast_admin/global"
	"github.com/labstack/gommon/color"
	"github.com/urfave/cli/v2"
	"net/url"
	"os/exec"
)

var EntMigrate = &cli.Command{
	Name:  "ent",
	Usage: "ent 数据生成及数据库同步",
	Action: func(ctx *cli.Context) error {
		//初始化配置文件
		bootstrap.ConfigInit()
		//fmt.Println("schema 代码生成")
		//cmd := exec.Command("go", "run", "-mod=mod", "entgo.io/ent/cmd/ent", "generate", "./ent/schema")
		//str, err := cmd.Output()
		//fmt.Println(str)
		//if err != nil {
		//	panic(err)
		//}
		var options []ent.Option
		if global.AppSetting.Database.Debug {
			options = append(options, ent.Debug())
		}
		client, err := ent.Open(
			"mysql",
			fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True&charset=utf8mb4&timeout=10s&readTimeout=60s&writeTimeout=60s&parseTime=True&loc=%s", global.AppSetting.Database.User, global.AppSetting.Database.Password, global.AppSetting.Database.Host, global.AppSetting.Database.Port, global.AppSetting.Database.Dbname, url.QueryEscape("Asia/Shanghai")),
			options...,
		)
		fmt.Println("数据库连接成功，正在同步")
		if err != nil {
			return errors.New(fmt.Sprintf("failed opening connection to sqlite: %v", err))
		}
		// 结构迁移
		err = client.Schema.Create(
			context.Background(),
			migrate.WithDropIndex(true),
			migrate.WithDropColumn(true),
		)
		if err != nil {
			return errors.New(fmt.Sprintf("failed migrate : %v", err))
		}
		client.Close()
		fmt.Println(color.Green("go_fast_admin: 同步完成,数据库连接关闭"))
		return nil
	},
	Subcommands: []*cli.Command{
		{
			Name:  "init",
			Usage: "ent 数据生成及数据库同步",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "classify",
					Aliases:     []string{"c"},
					Usage:       "生成类名",
					Required:    true,
				},
			},
			Action: func(ctx *cli.Context) error {
				name := ctx.String("classify")
				cmd := exec.Command("go", "run", "entgo.io/ent/cmd/ent", "init",name )
				_, err := cmd.Output()
				return err
			},
		},
	},
}
