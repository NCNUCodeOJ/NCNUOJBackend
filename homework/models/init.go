package models

import (
	//"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//DB 資料庫連接
var DB *gorm.DB

//Setup 資料庫連接設定
func Setup() {
	var err error
	DB, err = gorm.Open(sqlite.Open("test2.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	AutoMigrateAll()
}

//AutoMigrateAll 自動產生 table
func AutoMigrateAll() {
	DB.AutoMigrate(&Problem{})
	DB.AutoMigrate(&Tag{})
	DB.AutoMigrate(&Sample{})
	DB.AutoMigrate(&Tag2Table{})

}
