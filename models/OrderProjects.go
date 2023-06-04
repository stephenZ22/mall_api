package models

import (
	"time"

	"gorm.io/gorm"
)

// 订单产品连接表
type OrderProject struct {
	Id           uint `gorm:"unique_index;AUTO_INCREMENT;not null"`
	OrderID      uint `gorm:"primaryKey"`
	ProjectID    uint `gorm:"primaryKey"`
	ProjectCount uint

	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}
