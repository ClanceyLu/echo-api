package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ClanceyLu/echo-api/conf"
	"github.com/ClanceyLu/echo-api/pkg/logger"
	"github.com/ClanceyLu/echo-api/router"
)

func main() {
	// appConf app 基本配置
	appConf := conf.Conf.Sub("app")

	logger.Init()

	e := router.Init()
	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", appConf.GetInt("port")),
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}
	e.Logger.Fatal(e.StartServer(s))
}
