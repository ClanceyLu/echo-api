package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ClanceyLu/echo-api/conf"
	router "github.com/ClanceyLu/echo-api/routers"
)

func main() {
	// appConf app 基本配置
	appConf := conf.Conf.Sub("app")

	e := router.Init()
	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", appConf.GetInt("port")),
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}
	e.Logger.Fatal(e.StartServer(s))
}
