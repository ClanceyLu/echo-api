package custom

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// HTTPError 自定义 http error
type HTTPError struct {
	Code int    `json:"-"`
	Err  error  `json:"err"`
	Msg  string `json:"msg"`
}

func (he *HTTPError) Error() string {
	return he.Msg
}

// NewHTTPError 新建自定义HttpError
func NewHTTPError(code int, msg string) *HTTPError {
	return &HTTPError{
		Code: code,
		Msg:  msg,
	}
}

// NewParamError 新建参数错误
func NewParamError(err error) *HTTPError {
	return &HTTPError{
		422,
		err,
		err.Error(),
	}
}

// Error echo 错误集中处理
func Error(err error, c echo.Context) {
	var (
		code = http.StatusInternalServerError
		msg  = http.StatusText(code)
	)
	if he, ok := err.(*HTTPError); ok {
		msg = he.Msg
		log.Printf("HTTPError %+v", err)
		code = he.Code
	} else if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = http.StatusText(code)
	} else {
		msg = http.StatusText(code)
	}
	if !c.Response().Committed {
		r := Response{
			C:    c,
			Err:  err,
			Msg:  msg,
			Code: 1,
		}
		err := r.C.JSON(code, r)
		if err != nil {
			c.Logger().Error(err)
		}
	}
}
