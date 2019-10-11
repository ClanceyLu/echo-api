package router

import (
	"github.com/ClanceyLu/echo-api/controller"
	"github.com/ClanceyLu/echo-api/controller/admin"
	"github.com/labstack/echo/v4"
)

func adminRouter(r *echo.Group) {
	uploadController := controller.NewUpload()
	r.POST("/upload", uploadController.Upload)

	userController := admin.NewUser()
	r.GET("/user", userController.List)
}
