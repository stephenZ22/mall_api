package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserId  uint
	Status  uint
	Payment Payment
	Amount  uint `gorm:"type:decimal(50,2)"`

	Projects []Project `gorm:"many2many:order_projects;"`
}
