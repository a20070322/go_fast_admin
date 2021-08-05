package bootstrap

import (
	"fmt"
	"github.com/a20070322/go_fast_admin/global"
	"github.com/labstack/gommon/color"
	"github.com/spf13/viper"
)

func ConfigInit()  {
	v := viper.New()
	v.SetConfigName("config") // 指定配置文件
	v.SetConfigType("yaml")
	v.AddConfigPath("./")
	v.AddConfigPath("./config")
	err := v.ReadInConfig()        // 读取配置信息
	if err != nil {                // 读取配置信息失败
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	if err := v.Unmarshal(&global.AppSetting); err != nil {
		fmt.Printf("err:%s", err)
	}
	fmt.Println("go_fast_admin: "+color.Green("配置文件初始化成功"))
}