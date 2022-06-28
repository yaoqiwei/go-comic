package main

import (
	"fehu/common/lib"
	"fehu/router"
	"os"
	"os/signal"
	"syscall"
)

// @title API 文档
// @version 1.67
// @description  直播&点播 API文档.
// @host test.pgc.api.yimisaas.com
// @BasePath
func main() {
	lib.Init()

	defer lib.Destroy()

	router.HttpServerRun()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
