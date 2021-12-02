package models

import (
	"gorm.io/gorm"
)

// 把 TestPaper 裡面的 Name 改成 TestPaperName
//User Database - database
type TestPaper struct {
	gorm.Model
	TestPaperName string `gorm:"type:text;"`
	AuthorID      uint   `gorm:"NOT NULL;"`
	ClassID       uint   `gorm:"NOT NULL;"`
	Random        bool   `gorm:"NOT NULL;"`
	// 測驗卷名稱
	// 出卷者
	// 對應的課堂
	// 是否隨機出題
}

// AddTestPaper 新增測驗卷
func AddTestPaper(testpaper *TestPaper) {
	DB.Create(&testpaper)
}

// GetAllTestPapers 取得所有 testpaper
func GetAllTestPapers() (testpapers []TestPaper, err error) {
	err = DB.Find(&testpapers).Error
	return
}

// GetTestPaper 透過 ID 取得
func GetTestPaper(testpaperID uint) (TestPaper, error) {
	var testPaper TestPaper
	if err := DB.Where("id = ?", testpaperID).First(&testPaper).Error; err != nil {
		return TestPaper{}, err
	}
	return testPaper, nil
}

// EditTestPaper 修改
func EditTestPaper(testpaper *TestPaper) (err error) {
	err = DB.Where("id = ?", testpaper.ID).Save(&testpaper).Error
	return
}

// DeleteTestPaper 刪除
func DeleteTestPaper(testpaper TestPaper) (err error) {
	err = DB.Where("id = ?", testpaper.ID).Delete(&testpaper).Error
	return
}
