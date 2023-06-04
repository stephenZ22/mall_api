package models

import "gorm.io/gorm"

type UserInfo struct {
	gorm.Model
	UserId   int `gorm:"index:unique"`
	PhoneNum string
	Address  string
	Birthday string `gorm:"type:date"`
}
