package post

import (
	"github.com/ClanceyLu/echo-api/service"
	"github.com/labstack/echo/v4"
)

type post struct {
	service.Controller
}

type Post interface {
	PostPost(echo.Context) error
	GetPosts(echo.Context) error
}

func New(app service.Controller) Post {
	return &post{app}
}
