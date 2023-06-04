package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserInfo  UserInfo
	LoginName string `gorm:"type:varchar(50); unique_index"`
	PassWord  string `gorm:"type:varchar(50)"`

	Orders []Order
}
