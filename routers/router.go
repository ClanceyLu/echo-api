package router

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ClanceyLu/echo-api/conf"
	"github.com/ClanceyLu/echo-api/custom"
	v1 "github.com/ClanceyLu/echo-api/routers/v1"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Init 初始化 echo
func Init() *echo.Echo {
	appConf := conf.Conf.Sub("app")
	e := echo.New()
	e.Debug = appConf.GetBool("debug")
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// 自定义 Validate
	e.Validator = &custom.Validator{}
	// 自定义 Bind
	e.Binder = &custom.Binder{}

	// e.HTTPErrorHandler = handler.Error

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	e.POST("/register", v1.Register)

	e.PUT("/user/:id", v1.UpdateUser)

	// save router info to file
	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		log.Panic(err)
	}
	_ = ioutil.WriteFile("routes.json", data, 0644)

	return e
}
