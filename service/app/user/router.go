package user

import (
	"github.com/labstack/echo/v4"
)

type user struct{}

// User 是 user 接口定义
type User interface {
	GetUsers(echo.Context) error
	GetUser(echo.Context) error
}

// New 返回 User 接口实例
func New() User {
	return &user{}
}
