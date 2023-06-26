package models

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserInfo UserInfo `json:"-"`
	Email    string   `json:"email" gorm:"type:varchar(50); unique"`
	Password string   `gorm:"type:varchar(200)"`

	Orders []Order
}

func BcryptPassword(password string) (string, error) {
	bcryptPasswordByte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println("Error generating bcrypt hash:", err)
	}

	return string(bcryptPasswordByte), nil
}
