package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// IUpload 定义上传接口
type IUpload interface {
	Upload(c echo.Context) error
}

type uploadController struct{}

// NewUpload 返回 upload
func NewUpload() IUpload {
	return &uploadController{}
}

func (u uploadController) Upload(c echo.Context) error {
	return c.String(http.StatusOK, "upload")
}
