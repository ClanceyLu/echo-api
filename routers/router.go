package router

import (
	"net/http"

	"github.com/ClanceyLu/echo-api/conf"
	"github.com/ClanceyLu/echo-api/controller"
	"github.com/ClanceyLu/echo-api/custom"
	middle "github.com/ClanceyLu/echo-api/middleware"
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

	upload := controller.NewUpload()
	v1.POST("/upload", upload.Upload)

	// register app routers
	appRouter(v1.Group("/app"))
	// register admin routers
	adminRouter(v1.Group("/admin"))

	return e
}
