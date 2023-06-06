package models

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserInfo  UserInfo `json:"-"`
	LoginName string   `json:"login_name" gorm:"type:varchar(50); unique_index"`
	PassWord  string   `gorm:"type:varchar(200)"`

	Orders []Order
}

func BcryptPassword(password string) (string, error) {
	bcryptPasswordByte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println("Error generating bcrypt hash:", err)
	}

	return string(bcryptPasswordByte), nil
}
