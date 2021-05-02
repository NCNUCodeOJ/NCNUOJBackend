package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//DB 資料庫連接
var DB *gorm.DB

//Setup 資料庫連接設定
func Setup() {
	var err error
	DB, err = gorm.Open(sqlite.Open("testpaper.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	AutoMigrateAll()
}

//AutoMigrateAll 自動產生 table
func AutoMigrateAll() {
	DB.AutoMigrate(&Choice{})
	DB.AutoMigrate(&ChoiceOption{})
	DB.AutoMigrate(&ChoiceTP{})
	DB.AutoMigrate(&Cloze{})
	DB.AutoMigrate(&ClozeAnswer{})
	DB.AutoMigrate(&ClozeTP{})
	DB.AutoMigrate(&TestPaper{})
}
