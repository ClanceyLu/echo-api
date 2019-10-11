package router

import (
	"github.com/ClanceyLu/echo-api/controller"
	"github.com/ClanceyLu/echo-api/controller/app"
	"github.com/labstack/echo/v4"
)

func appRouter(r *echo.Group) {
	uploadController := controller.NewUpload()
	r.POST("/upload", uploadController.Upload)

	userController := app.NewUser()
	r.GET("/user", userController.List)
}
