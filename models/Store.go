package models

import "gorm.io/gorm"

type Store struct {
	gorm.Model
	UserId uint
	User   User

	Projects []Project
	Name     string
	Status   uint
	StoreNum uint
}
