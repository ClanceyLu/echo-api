package admin

import (
	"github.com/ClanceyLu/echo-api/custom"
	"github.com/ClanceyLu/echo-api/models"
	"github.com/labstack/echo/v4"
)

// IUser 定义 admin 的用户接口
type IUser interface {
	List(c echo.Context) error
}

type userController struct {
	model models.IUser
}

// NewUser 返回用户接口
func NewUser() IUser {
	return &userController{
		model: models.NewUserModel(),
	}
}

func (u *userController) List(c echo.Context) error {
	cc := c.(*custom.Context)
	user, err := u.model.GetDetail(1)
	if err != nil {
		return custom.NewHTTPError(404, "啊啊啊").
			SetErr(err)
	}
	return cc.Suc(user)
}
