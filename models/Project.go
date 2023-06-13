package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	StoreId uint    `json:"store_id"`
	Store   Store   `json:"-"`
	Name    string  `json:"name"`
	Price   float32 `gorm:"type:decimal(9,2)" json:"price"`
	Status  uint    `gorm:"default:0" json:"status"`
}
