package db

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MainDb *gorm.DB

func DataBaseConnection() {
	var err error

	dsn := os.Getenv("DataBaseDsn")
	MainDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to Database")
	}

}
