package v1

import (
	"log"
	"strconv"

	"github.com/ClanceyLu/echo-api/custom"
	"github.com/ClanceyLu/echo-api/models"
	"github.com/ClanceyLu/echo-api/service"
	"github.com/labstack/echo/v4"
)

// Register 用户注册
func Register(c echo.Context) error {
	r := custom.Response{C: c}
	var user models.User

	if err := c.Bind(&user); err != nil {
		return err
	}
	log.Printf("user %+v", user)

	if err := service.Register(&user); err != nil {
		return err
	}

	return r.Response(nil)
}

// UpdateUser 跟新用户
func UpdateUser(c echo.Context) error {
	r := custom.Response{C: c}
	id, _ := strconv.Atoi(c.Param("id"))
	var user models.User
	if err := c.Bind(&user); err != nil {
		return err
	}
	user.ID = uint(id)
	if err := service.UpdateUser(&user); err != nil {
		return err
	}
	return r.Response(user)
}
