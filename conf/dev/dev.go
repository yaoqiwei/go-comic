package dev

import (
	. "fehu/conf/structs"
	"io/ioutil"
	"time"

	"github.com/sirupsen/logrus"
)

var Base BaseConf = BaseConf{
	DebugMode: "debug",
	TimeZone:  time.FixedZone("CST", 8*3600),
	Http: HttpConf{
		Addr:         ":20150",
		UploadDomain: "https://test.pgc.img.yimisaas.com",
		UploadExec:   "/u.php",
		UploadAuth:   "A958EFDDB469B74FF3DEB81717914225",
		ApiAesKey:    "C77A9FF7F323B5404902102257503C2F",
	},

	Log: LogConf{
		Level:      logrus.TraceLevel,
		TimeLayout: "2006/01/02 - 15:04:05.000",
	},

	Redis: map[string]*RedisConf{
		"default": {
			ProxyList: []string{"127.0.0.1:6379"},
			MaxActive: 100,
			MaxIdle:   20,
		},
	},

	Mysql: MysqlMapConfig{
		//List: map[string]*MysqlConf{
		//	"default": {
		//		DriverName:      "mysql",
		//		DataSourceName:  "root:admin0805@tcp(192.168.24.214:3306)/yl_comic?charset=utf8mb4&parseTime=true&loc=Asia%2FChongqing",
		//		MaxOpenConn:     50,
		//		MaxIdleConn:     20,
		//		MaxConnLifeTime: 100,
		//		Prefix:          "comic_",
		//	},
		//	"read": {
		//		DriverName:      "mysql",
		//		DataSourceName:  "root:admin0805@tcp(192.168.24.214:3306)/yl_comic?charset=utf8mb4&parseTime=true&loc=Asia%2FChongqing",
		//		MaxOpenConn:     50,
		//		MaxIdleConn:     20,
		//		MaxConnLifeTime: 100,
		//		Prefix:          "comic_",
		//	},
		//},
		List: map[string]*MysqlConf{
			"default": {
				DriverName:      "mysql",
				DataSourceName:  "root:960216@tcp(101.43.63.154:3306)/yl_comic?charset=utf8mb4&parseTime=true&loc=Asia%2FChongqing",
				MaxOpenConn:     50,
				MaxIdleConn:     20,
				MaxConnLifeTime: 100,
				Prefix:          "comic_",
			},
			"read": {
				DriverName:      "mysql",
				DataSourceName:  "root:960216@tcp(101.43.63.154:3306)/yl_comic?charset=utf8mb4&parseTime=true&loc=Asia%2FChongqing",
				MaxOpenConn:     50,
				MaxIdleConn:     20,
				MaxConnLifeTime: 100,
				Prefix:          "comic_",
			},
		},
		Split: 4,
	},

	WechatApi: WechatApiConf{},

	Swag: SwagConf{
		Enable:      true,
		Version:     GetVersion(),
		Host:        "127.0.0.1:20150",
		BasePath:    "/",
		Schemes:     []string{"http"},
		Title:       "API 文档",
		Description: "直播&点播 API文档.",
		Url:         "http://127.0.0.1:20150/swagger/doc.json",
	},

	Es: EsConf{
		Address:  "http://47.114.93.122:9200",
		UserName: "elastic",
		Password: "49e40e6da1645eff828ba76f35b7c31b",
	},
}

func GetVersion() string {
	content, _ := ioutil.ReadFile("version")
	return string(content)
}
