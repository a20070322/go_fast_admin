package types

type AppConfigure struct {
	Version   string           `yaml:"Version"`
	Server    *ServerConfig    `yaml:"Server"`
	Jwt       *JwtConfig       `yaml:ConfigJwt`
	Redis     *RedisConfig     `yaml:Redis`
	Database  *DatabaseConfig  `yaml:Database`
	Logger    *LoggerConfig    `yaml:Logger`
	Env       string           `yaml:Env`
	MinChat   *MinChatConfig   `yaml:MinChat`
	RabbitMq  *RabbitMqConfig  `yaml:RabbitMq`
	WechatPay *WechatPayConfig `yaml:WechatPay`
}

type JwtConfig struct {
	JwtSecret           string `yaml:"jwtSecret"`
	TokenExpireDuration int    `yaml:"tokenExpireDuration"`
}

type ServerConfig struct {
	Port  int  `yaml:"port"`
	Debug bool `yaml:"debug"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type DatabaseConfig struct {
	Debug    bool   `yaml:"debug"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Migrate  bool   `yaml:"migrate"`
}
type LoggerConfig struct {
	LogPath    string `yaml:"logPath"`
	LogErrPath string `yaml:"logErrPath"`
	Level      string `yaml:"level"`
}

type MinChatConfig struct {
	AppId  string `yaml:"appId"`
	Secret string `yaml:"secret"`
}

type RabbitMqConfig struct {
	UserName string `yaml:"userName"`
	PassWord string `yaml:"passWord"`
	IP       string `yaml:"ip"`
	Port     int    `yaml:"port"`
	VHost    string `yaml:"vHost"`
}

type WechatPayConfig struct {
	MchID                      string `yaml:"MchID"`
	MchCertificateSerialNumber string `yaml:"MchCertificateSerialNumber"`
	MchPrivateKey              string `yaml:"MchPrivateKey"`
	MchAPIv3Key                string `yaml:"MchAPIv3Key"`
}

const (
	EnvLOCAL  = "local"
	EnvDEV    = "dev"
	EnvONLINE = "online"
)
