package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// IUser represents controller of user
type IUser interface {
	List(c echo.Context) error
}

type userController struct{}

// NewUser return an user controller
func NewUser() IUser {
	return &userController{}
}

func (u userController) List(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}
