package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	StoreId uint
	Store   Store
	Name    string
	Price   uint `gorm:"type:decimal(9,2)"`
	Status  uint `gorm:"default:0"`
}
