package models

import "gorm.io/gorm"

type Store struct {
	gorm.Model
	UserId uint
	User   User

	Projects []Project
	Status   uint
	StoreNum uint
}
