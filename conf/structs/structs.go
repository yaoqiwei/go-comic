package structs

import (
	"time"

	"github.com/sirupsen/logrus"
)

type HttpConf struct {
	TimeZone       *time.Location
	Addr           string
	AesOpen        bool
	ApiAesKey      string
	HeaderCheck    bool
	UploadDomain   string
	UploadExec     string
	UploadAuth     string
	TrustedProxies []string
}

type RedisConf struct {
	ProxyList      []string
	MaxActive      int // 最大活动连接数，值为0时表示不限制
	MaxIdle        int // 最大空闲连接数
	Downgrade      bool
	Network        string // 通讯协议，默认为 tcp
	Password       string // redis鉴权密码
	Db             int    // 数据库
	IdleTimeout    int    // 空闲连接的超时时间，超过该时间则关闭连接。单位为秒。默认值是5分钟。值为0时表示不关闭空闲连接。此值应该总是大于redis服务的超时时间。
	Prefix         string // 键名前缀
	ConnectTimeout int
	ReadTimeout    int
	WriteTimeout   int
	Wait           bool
}

type MysqlConf struct {
	DriverName      string
	DataSourceName  string
	MaxOpenConn     int
	MaxIdleConn     int
	MaxConnLifeTime int
	Prefix          string
}

type MysqlMapConfig struct {
	List  map[string]*MysqlConf
	Split int
}

type LogConf struct {
	Level      logrus.Level
	TimeLayout string
}

type WechatApiConf struct {
	AppId                      string
	AppSecret                  string
	OffiAppId                  string
	OffiAppSecret              string
	OffiMessageTemplate        string
	Mchid                      string
	NotifyUrl                  string
	MchCertificateSerialNumber string
	MchAPIv3Key                string
}

type SwagConf struct {
	Enable      bool
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
	Url         string
}

type EsConf struct {
	Address  string
	UserName string
	Password string
}

type BaseConf struct {
	Http      HttpConf
	TimeZone  *time.Location
	DebugMode string
	Redis     map[string]*RedisConf
	Mysql     MysqlMapConfig
	Log       LogConf
	WechatApi WechatApiConf
	Swag      SwagConf
	Es        EsConf
}
