package admin

import (
	"net/http"
	"strconv"

	"github.com/ClanceyLu/echo-api/models"
	"github.com/labstack/echo/v4"
)

// IUser 定义 admin 的用户接口
type IUser interface {
	List(c echo.Context) error
	Add(c echo.Context) error
	Detail(c echo.Context) error
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
	return c.String(http.StatusOK, "")
}

func (u *userController) Add(c echo.Context) error {
	user := &models.User{}
	if err := c.Bind(user); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

// Detail 根据 ID 返回用户详情
func (u *userController) Detail(c echo.Context) error {
	id := c.Param("id")
	uID, _ := strconv.Atoi(id)

	// 判断用户是否存在
	if exist, err := u.model.ExistByID(uint(uID)); err != nil {
		return err
	} else if !exist {
		return echo.NewHTTPError(http.StatusNotFound, "用户不存在")
	}

	user, err := u.model.GetDetail(uID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}
