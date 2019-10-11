package models

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

// User 用户表
type User struct {
	Model
	User     string  `gorm:"comment:'账号'" json:"user" valid:"required~用户名不能为空"`
	Password string  `gorm:"comment:'密码'" json:"password" valid:"required~密码不能为空,stringlength(8|20)~密码必须是8-20位"`
	NickName string  `gorm:"comment:'用户名'" json:"nickName" valid:"omitempty,stringlength(4|10)"`
	PhoneNo  string  `gorm:"comment:'手机号'" json:"phoneNo" valid:"required"`
	Email    string  `gorm:"comment:'邮箱'" json:"email" valid:"email"`
	Sex      uint    `gorm:"default:3;comment:'性别，1:男 2:女 3:未知'" valid:"required,in(0|1|2)" json:"sex"`
	Profile  Profile `gorm:"type:json" json:"profile"`
}

// Profile 用户信息
type Profile struct {
	Email   string `json:"email" valid:"email"`
	PhoneNo string `json:"phoneNo"`
}

// Value 实现 JSON 接口
func (p Profile) Value() (driver.Value, error) {
	b, err := json.Marshal(p)
	return string(b), err
}

// Scan 实现 JSON 接口
func (p *Profile) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), p)
}

// IUser is user model interface
type IUser interface {
	Add(u *User) error
	Update(id uint, u *User) error
	ExistByPhoneNo(phoneNo string) (bool, error)
	ExistByID(ID uint) (bool, error)
	GetDetail(query interface{}) (*User, error)
}

type userModel struct{}

// NewUserModel return IUser
func NewUserModel() IUser {
	return &userModel{}
}

func (u *User) save(tx *gorm.DB) error {
	return tx.Save(u).Error
}

// Add 新增用户
func (um *userModel) Add(u *User) error {
	return u.save(db)
}

// Update 更新用户
func (um *userModel) Update(id uint, u *User) error {
	return db.
		Model(&User{}).
		Where("id = ?", id).
		Updates(u).
		Error
}

// ExistByPhoneNo 用户手机是否已经注册过
func (um *userModel) ExistByPhoneNo(phoneNo string) (bool, error) {
	var u User
	if err := db.
		Select("id").
		Where("phone_no = ?", phoneNo).
		First(&u).
		Error; dbErr(err) {
		return false, err
	}
	if u.ID > 0 {
		return true, nil
	}
	return false, nil
}

// ExistByID 用户手机是否已经注册过
func (um *userModel) ExistByID(ID uint) (bool, error) {
	var u User
	if err := db.
		Select("id").
		Where("id = ?", ID).
		First(&u).
		Error; dbErr(err) {
		return false, err
	}
	if u.ID > 0 {
		return true, nil
	}
	return false, nil
}

// GetDetail 根据查询条件返回用户详情
func (um *userModel) GetDetail(query interface{}) (*User, error) {
	var u User
	if err := db.
		Where(query).
		First(&u).
		Error; err != nil {
		return nil, errors.Wrapf(err, "user %+v not found", query)
	}
	return &u, nil
}
