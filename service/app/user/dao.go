package user

import (
	"github.com/ClanceyLu/echo-api/model"
)

type listQuery struct {
	Page     uint
	PageSize uint
	Where    map[string]interface{}
}

func (u *userService) queryUsers(query *listQuery) (*[]model.User, int, error) {
	var users []model.User
	db := u.db
	db = db.Model(&model.User{}).
		Select("id, nick_name, phone_no, email")
	db = db.Where(query.Where)
	offset := (query.Page - 1) * query.PageSize
	count := 0
	db = db.
		Count(&count).
		Offset(offset).
		Limit(query.PageSize)
	if err := db.
		Find(&users).
		Error; err != nil {
		return nil, 0, err
	}
	return &users, count, nil
}

func (u *userService) queryUserByID(id uint) (*model.User, error) {
	db := u.db
	user := &model.User{}
	if err := db.
		Select("id, nick_name, user, phone_no, email, sex, profile").
		First(user, id).
		Error; err != nil {
		return nil, err
	}
	return user, nil
}

// 判断手机号是否注册过
func (u *userService) existByPhoneNo(phoneNo string) (bool, error) {
	user := &model.User{}
	if err := u.db.
		Select("id").
		Where("phone_no = ?", phoneNo).
		First(u).
		Error; err != nil {
		return false, err
	}
	if user.ID > 0 {
		return true, nil
	}
	return false, nil
}
