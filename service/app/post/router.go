package post

import (
	"github.com/labstack/echo/v4"
)

type post struct{}

// Post 接口定义
type Post interface {
	PostPost(echo.Context) error
	GetPosts(echo.Context) error
}

// New 返回 Post 接口实例
func New() Post {
	return &post{}
}
