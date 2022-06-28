package lib

import (
	"fehu/conf"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
)

// Init 模块初始化
func Init() error {

	time.Local = conf.TimeZone
	color.NoColor = false

	logger := logrus.StandardLogger()
	logger.Out = os.Stdout

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		ForceColors:     true,
		TimestampFormat: conf.Log.TimeLayout,
	})
	logrus.SetLevel(conf.Log.Level)

	if err := InitRedis(); err != nil {
		logrus.Errorf("InitRedisConf:" + err.Error())
	}

	if err := InitDBPool(); err != nil {
		logrus.Errorf("InitDBPool:" + err.Error())
	}
	if err := InitGormPool(); err != nil {
		logrus.Errorf("InitGromPool:" + err.Error())
	}

	logrus.Infof("success loading resources.")
	logrus.Infof("------------------------------------------------------------------------")
	return nil
}

// Destroy 公共销毁函数
func Destroy() {
	logrus.Infof("------------------------------------------------------------------------")
	logrus.Infof("start destroy resources.")
	CloseDB()
	logrus.Infof("success destroy resources.")
}
