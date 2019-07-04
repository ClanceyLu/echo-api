package custom

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Response 自定义返回结构
type Response struct {
	Code int          `json:"code"`
	Err  error        `json:"err"`
	Msg  string       `json:"msg"`
	Data interface{}  `json:"data"`
	C    echo.Context `json:"-"`
}

// Response Response Suc handler function
func (r *Response) Response(data interface{}) error {
	r.Code = 0
	r.Data = data
	return r.C.JSON(http.StatusOK, data)
}

// ErrResponse Response Err handler function
func (r *Response) ErrResponse(msg string) error {
	r.Code = 1
	r.Msg = msg
	return r.C.JSON(http.StatusOK, nil)
}

// ServerErr Response Server Error
func (r *Response) ServerErr(err error) error {
	r.Err = err
	return r.C.JSON(http.StatusInternalServerError, nil)
}
