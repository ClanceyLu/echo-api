package post

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (p postService) postPost(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}
