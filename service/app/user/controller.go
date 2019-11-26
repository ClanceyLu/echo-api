package user

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (u user) GetUser(c echo.Context) error {
	user, err := u.queryUserByID(1)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, echo.Map{
		"user": user,
	})
}

func (u user) GetUsers(c echo.Context) error {
	query := listQuery{
		Page:     1,
		PageSize: 10,
	}
	list, count, err := u.queryUsers(&query)
	if err != nil {
		return err
	}
	log.Printf("list %+v", list)
	return c.JSON(http.StatusOK, echo.Map{
		"list":  list,
		"count": count,
	})
}
