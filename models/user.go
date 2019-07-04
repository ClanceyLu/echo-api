package models

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/jinzhu/gorm"
)

func dbErr(err error) bool {
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		return true
	}
	return false
}

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

func (u *User) save(tx *gorm.DB) error {
	return tx.Save(u).Error
}

// AddUser 新增用户
func AddUser(u *User) error {
	return u.save(db)
}

// UpdateUser 更新用户
func UpdateUser(id uint, u *User) error {
	return db.
		Model(&User{}).
		Where("id = ?", id).
		Update(u).
		Error
}

// ExistUserByPhoneNo 用户手机是否已经注册过
func ExistUserByPhoneNo(phoneNo string) (bool, error) {
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
