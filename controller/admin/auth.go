package admin

import (
	"net/http"
	"time"

	"github.com/ClanceyLu/echo-api/conf"
	"github.com/ClanceyLu/echo-api/models"
	"github.com/ClanceyLu/echo-api/pkg/auth"
	"github.com/labstack/echo/v4"
)

// IAuth admin 鉴权接口
type IAuth interface {
	Login(echo.Context) error
	Register(echo.Context) error
}

type authController struct{}

// NewAuth return authController instance
func NewAuth() IAuth {
	return &authController{}
}

// Login 是 admin 用户登陆
func (a authController) Login(c echo.Context) error {
	type body struct {
		UserName string `valid:"required~用户名不能为空"`
		Password string `valid:"required~密码不能为空,stringlength(8|20)~密码不正确"`
	}
	authBody := &body{}
	if err := c.Bind(authBody); err != nil {
		return err
	}

	// 查找对应用户
	query := make(map[string]interface{})
	query["userName"] = authBody.UserName
	authConf := conf.Conf.Sub("auth")
	salt := authConf.GetString("salt")
	cryptoPassword := auth.CryptoPassword(authBody.Password, salt)
	query["password"] = cryptoPassword
	userModel := models.NewUserModel()
	user, err := userModel.GetDetail(query)
	if err != nil {
		return err
	}
	if user.ID == 0 {
		return echo.NewHTTPError(http.StatusOK, "用户名或密码不正确")
	}

	// 生成 token 返回信息
	payload := map[string]interface{}{
		"ID":      user.ID,
		"PhoneNo": user.PhoneNo,
		"exp":     time.Now().AddDate(0, 1, 0).Unix(),
	}
	secret := authConf.GetString("secret")
	token, err := auth.CryptoToken(payload, secret)
	return c.JSON(http.StatusOK, echo.Map{
		"user":  user,
		"token": token,
	})
}

func (a authController) Register(c echo.Context) error {
	return nil
}
