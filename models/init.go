package models

import (
	"startlab/ChartOnWeb/dataaccess"

	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dataaccess.InitializeDatabase()
	DB = dataaccess.DB
}
