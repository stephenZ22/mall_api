package db

import "MallApi/models"

func DataBaseMigrate() {
	MainDb.AutoMigrate(&models.User{})
	MainDb.AutoMigrate(&models.UserInfo{})
	MainDb.AutoMigrate(&models.Store{})
	MainDb.AutoMigrate(&models.Project{})
	MainDb.AutoMigrate(&models.Order{})
	MainDb.AutoMigrate(&models.OrderProject{})
	MainDb.AutoMigrate(&models.Payment{})

}
