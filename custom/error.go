package custom

import (
	"log"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

// Error echo 错误集中处理
func Error(err error, c echo.Context) {
	var (
		status = http.StatusInternalServerError
		msg    = http.StatusText(status)
	)
	log.Printf("catch err %+v", err)
	if err == gorm.ErrRecordNotFound {
		// mysql not found resource, return 404
		status = http.StatusNotFound
		msg = http.StatusText(status)
	} else if e, ok := err.(govalidator.Error); ok {
		log.Print("catch validate err")
		msg = e.Error()
		status = http.StatusUnprocessableEntity
	} else if he, ok := err.(*echo.HTTPError); ok {
		status = he.Code
		msg = http.StatusText(status)
	} else {
		msg = http.StatusText(status)
	}
	if !c.Response().Committed {
		r := Response{
			C:    c,
			Err:  err,
			Msg:  msg,
			Code: 10001,
		}
		// not show error to front end if app is not debug mode
		if !c.Echo().Debug {
			r.Err = nil
		}
		err := r.C.JSON(status, r)
		if err != nil {
			c.Logger().Error(err)
		}
	}
}
