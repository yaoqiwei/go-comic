package conf

import (
	"fehu/conf/dev"
	structs "fehu/conf/structs"
	"fehu/conf/test"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var Base structs.BaseConf
var Http structs.HttpConf
var DebugMode string
var TimeZone *time.Location
var Redis map[string]*structs.RedisConf
var Mysql structs.MysqlMapConfig
var Log structs.LogConf
var WechatApi structs.WechatApiConf
var Swag structs.SwagConf
var Es structs.EsConf

// GetEnvironment 获取环境变量
func GetEnvironment() string {
	content, _ := ioutil.ReadFile("environment")
	env := string(content)
	if env == "" {
		env = os.Getenv("ENVIRONMENT")
	}
	return env
}

// init 初始化godotenv组件
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("no such .env file, use system env")
	}
}

// init 初始化配置选择
func init() {
	switch GetEnvironment() {
	case "dev":
		Base = dev.Base
	case "test":
		Base = test.Base
	default:
		Base = dev.Base
	}

	Http = Base.Http
	DebugMode = Base.DebugMode
	TimeZone = Base.TimeZone
	Redis = Base.Redis
	Mysql = Base.Mysql
	Log = Base.Log
	WechatApi = Base.WechatApi
	Swag = Base.Swag
	Es = Base.Es
}
