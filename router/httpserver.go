package router

import (
	"fehu/conf"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func HttpServerRun() {
	gin.SetMode("release")
	r := InitRouter()

	addr := conf.Http.Addr
	go func() {
		logrus.Infof("HttpServerRun:%s", addr)
		if err := r.Run(addr); err != nil {
			logrus.Errorf("HttpServerRun:%s err:%v", addr, err)
		}
	}()
}
