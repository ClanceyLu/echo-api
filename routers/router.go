package router

import (
	"log"
	"net/http"

	"github.com/ClanceyLu/echo-api/conf"
	"github.com/ClanceyLu/echo-api/controller"
	"github.com/ClanceyLu/echo-api/controller/admin"
	"github.com/ClanceyLu/echo-api/controller/app"
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

	e.GET("/ping", func(c echo.Context) error {
		cc := c.(*custom.Context)
		arr := cc.QueryArray("aa")
		log.Print(arr)
		page, pageSize := cc.PageInfo()
		log.Printf("page %d, pageSize %d", page, pageSize)
		return c.String(http.StatusOK, "pong")
	})

	uploadController := controller.NewUpload()
	appRouter := v1.Group("/app")
	{
		appRouter.POST("/upload", uploadController.Upload)

		userController := app.NewUser()
		appRouter.GET("/user", userController.List)
	}

	adminRouter := v1.Group("/admin")
	{
		adminRouter.POST("/upload", uploadController.Upload)

		userController := admin.NewUser()
		adminRouter.GET("/user", userController.List)
	}

	return e
}
