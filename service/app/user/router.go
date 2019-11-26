package user

import (
	"github.com/ClanceyLu/echo-api/service"
	"github.com/labstack/echo/v4"
)

type user struct {
	service.Controller
}

// User 是 user 接口定义
type User interface {
	GetUsers(echo.Context) error
	GetUser(echo.Context) error
}

// New 返回 UserService
func New(app service.Controller) User {
	return &user{app}
}
