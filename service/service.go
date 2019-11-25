package service

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

// Service 定义了服务接口
type Service interface {
	Router(r *echo.Group)
}

// Controller 定义了 controller 的结构
type Controller struct {
	DB     *gorm.DB
	Router *echo.Group
}
