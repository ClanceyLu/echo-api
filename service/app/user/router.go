package user

import (
	"github.com/ClanceyLu/echo-api/service"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type userService struct {
	db *gorm.DB
}

// New 返回 UserService
func New(db *gorm.DB) service.Service {
	return &userService{db}
}

// Router 注册路由
func (u *userService) Router(r *echo.Group) {
	r.GET("/user", u.getUsers)
	r.GET("/user/:id", u.getUser)
}
