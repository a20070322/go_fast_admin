package bootstrap

import (
	"context"
	"fmt"
	"github.com/a20070322/go_fast_admin/global"
	"github.com/go-redis/redis/v8"
	"github.com/labstack/gommon/color"
)

func RedisConfig() {
	ctx := context.Background()
	global.Rdb = redis.NewClient(&redis.Options{
		Addr:    fmt.Sprintf("%s:%d",global.AppSetting.Redis.Host,global.AppSetting.Redis.Port),
		Password: global.AppSetting.Redis.Password, // no password set
		DB:       global.AppSetting.Redis.DB,          // use default DB
	})
	_, err := global.Rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Redis is error")
		panic(err.Error())
	}
	fmt.Println("go_fast_admin: "+color.Green("Redis初始化成功"))
}
