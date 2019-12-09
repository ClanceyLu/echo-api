package upload

import (
	"github.com/ClanceyLu/echo-api/service"
	"github.com/labstack/echo/v4"
)

type uploadService struct{}

// New 返回一个 upload 服务
func New() service.Service {
	return &uploadService{}
}

func (u *uploadService) Router(r *echo.Group) {
	r.POST("/upload", u.upload)
}
