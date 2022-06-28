package test

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
			ProxyList: []string{"172.20.0.1:6379"},
			MaxActive: 100,
			MaxIdle:   20,
			Password:  "woaini123",
		},
	},

	Mysql: MysqlMapConfig{
		List: map[string]*MysqlConf{
			"default": {
				DriverName:      "mysql",
				DataSourceName:  "thinkcmf:MweYpMEAFni5n44k@tcp(172.20.0.1:3306)/thinkcmf?charset=utf8mb4&parseTime=true&loc=Asia%2FChongqing",
				MaxOpenConn:     50,
				MaxIdleConn:     20,
				MaxConnLifeTime: 100,
				Prefix:          "cmf_",
			},
			"read": {
				DriverName:      "mysql",
				DataSourceName:  "thinkcmf:MweYpMEAFni5n44k@tcp(172.20.0.1:3306)/thinkcmf?charset=utf8mb4&parseTime=true&loc=Asia%2FChongqing",
				MaxOpenConn:     50,
				MaxIdleConn:     20,
				MaxConnLifeTime: 100,
				Prefix:          "cmf_",
			},
		},
		Split: 3,
	},

	WechatApi: WechatApiConf{},

	Swag: SwagConf{
		Enable:      true,
		Version:     GetVersion(),
		Host:        "test.pgc.api.yimisaas.com",
		BasePath:    "/",
		Schemes:     []string{"http"},
		Title:       "API 文档",
		Description: "直播&点播 API文档.",
		Url:         "http://test.pgc.api.yimisaas.com/swagger/doc.json",
	},

	Es: EsConf{
		Address:  "http://172.20.0.1:9200",
		UserName: "elastic",
		Password: "49e40e6da1645eff828ba76f35b7c31b",
	},
}

func GetVersion() string {
	content, _ := ioutil.ReadFile("version")
	return string(content)
}
