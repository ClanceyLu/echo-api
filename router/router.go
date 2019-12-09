package router

import (
	"net/http"

	"github.com/ClanceyLu/echo-api/conf"
	"github.com/ClanceyLu/echo-api/custom"
	middle "github.com/ClanceyLu/echo-api/middleware"
	"github.com/ClanceyLu/echo-api/service/app"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Init 初始化 echo
func Init() *echo.Echo {
	appConf := conf.Conf.Sub("app")
	e := echo.New()
	e.Debug = appConf.GetBool("debug")
	e.Use(middle.CustomContext())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// 自定义 Validate
	e.Validator = &custom.Validator{}
	// 自定义 Bind
	e.Binder = &custom.Binder{}

	// custom errorHandler
	e.HTTPErrorHandler = custom.Error

	v1 := e.Group("/v1")
	{
		v1.GET("/ping", func(c echo.Context) error {
			return c.String(http.StatusOK, "pong")
		})
	}

	// register app routers
	app := app.New()
	app.Router(v1)

	return e
}
