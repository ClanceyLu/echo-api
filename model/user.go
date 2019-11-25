package model

import (
	"database/sql/driver"
	"encoding/json"
)

// User 用户表
type User struct {
	Model
	User     string   `gorm:"comment:'账号'" json:"user,omitempty" valid:"required~用户名不能为空"`
	Password string   `gorm:"comment:'密码'" json:"password,omitempty" valid:"required~密码不能为空,stringlength(8|20)~密码必须是8-20位"`
	NickName string   `gorm:"comment:'用户名'" json:"nickName,omitemty" valid:"omitempty,stringlength(4|10)"`
	PhoneNo  string   `gorm:"comment:'手机号'" json:"phoneNo,omitemty" valid:"required"`
	Email    string   `gorm:"comment:'邮箱'" json:"email,omitemty" valid:"email"`
	Sex      uint     `gorm:"default:3;comment:'性别，1:男 2:女 3:未知'" valid:"required,in(0|1|2)" json:"sex,omitempty"`
	Profile  *Profile `gorm:"type:json" json:"profile,omitempty"`
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
