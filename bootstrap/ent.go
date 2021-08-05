package bootstrap

import (
	"context"
	"fmt"
	"github.com/a20070322/go_fast_admin/ent"
	"github.com/a20070322/go_fast_admin/ent/migrate"
	"github.com/a20070322/go_fast_admin/global"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/gommon/color"
	"net/url"
	"time"
)

func EntInit() {
	var options []ent.Option
	if global.AppSetting.Database.Debug {
		options = append(options, ent.Debug())
	}
	client, err := ent.Open(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True&charset=utf8mb4&timeout=10s&readTimeout=60s&writeTimeout=60s&parseTime=True&loc=%s", global.AppSetting.Database.User, global.AppSetting.Database.Password, global.AppSetting.Database.Host, global.AppSetting.Database.Port, global.AppSetting.Database.Dbname,url.QueryEscape("Asia/Shanghai")),
		options...,
	)
	if err != nil {
		panic(fmt.Sprintf("failed opening connection to sqlite: %v", err))
	}
	ctx := context.Background()

	// 结构迁移
	if global.AppSetting.Database.Migrate == true {
		err = client.Schema.Create(
			ctx,
			migrate.WithDropIndex(true),
			migrate.WithDropColumn(true),
		)
	}
	if err != nil {
		panic(fmt.Sprintf("failed creating schema resources: %v", err))
	}

	//记录执行时间  目前只针对插入修改删除生效
	client.Use(func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			start := time.Now()
			defer func() {
				global.Logger.Info(fmt.Sprintf("Op=%s\tType=%s\tTime=%s\tConcreteType=%T\n", m.Op(), m.Type(), time.Since(start), m))
			}()
			return next.Mutate(ctx, m)
		})
	})

	//记录错误日志 目前只针对插入修改删除生效
	client.Use(func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			v, err := next.Mutate(ctx, m)
			if err != nil {
				global.Logger.Error(err.Error())
				return nil, err
			}
			return v, nil
		})
	})
	global.Db = client

	//err = casbin_rules_service.Init(context.Background()).InitCasbinEnt()
	//fmt.Println(err)
	fmt.Println("go_fast_admin: "+color.Green("数据库初始化成功"))
}