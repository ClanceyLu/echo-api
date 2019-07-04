package service

import (
	"github.com/ClanceyLu/echo-api/custom"
	"github.com/ClanceyLu/echo-api/models"
)

// Register 注册用户
func Register(u *models.User) error {
	// 手机号是否注册过
	if exist, err := models.ExistUserByPhoneNo(u.PhoneNo); err != nil {
		return err
	} else if exist {
		return custom.NewHTTPError(200, "手机号已经注册过")
	}

	// 创建用户
	return models.AddUser(u)
}

// UpdateUser 更新用户信息
func UpdateUser(u *models.User) error {
	return models.UpdateUser(u.ID, u)
}
