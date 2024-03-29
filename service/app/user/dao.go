package user

import (
	"github.com/ClanceyLu/echo-api/model"
	"github.com/ClanceyLu/echo-api/service"
)

type listQuery struct {
	Page     uint
	PageSize uint
	Where    map[string]interface{}
}

func (u user) queryUsers(query *listQuery) (*[]model.User, int, error) {
	var users []model.User
	db := service.Mysql
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

func (u user) queryUserByID(id uint) (*model.User, error) {
	db := service.Mysql
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
func (u user) existByPhoneNo(phoneNo string) (bool, error) {
	user := &model.User{}
	if err := service.Mysql.
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
