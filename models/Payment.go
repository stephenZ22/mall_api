package models

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	OrderId uint
	Amount  uint `gorm:"type:decimal(50,2)"`
}
