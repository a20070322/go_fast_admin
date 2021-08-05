package bootstrap

func Init() {
	//配置文件初始化
	ConfigInit()
	//日志初始化
	LoggerInit()
	// redis初始化
	RedisConfig()
	//数据库初始化
	EntInit()
	//权限初始化
	CasbinInit()


}
