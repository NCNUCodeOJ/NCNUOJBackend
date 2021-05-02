package models

import "gorm.io/gorm"

//User Database - database
type TestPaper struct {
	gorm.Model
	Name     string `gorm:"NOT NULL;"`
	AuthorID uint
	ClassID  uint
	Random   bool
	// 測驗卷名稱
	// 出卷者
	// 對應的課堂
	// 是否隨機出題
}

// AddTestPaper 新增測驗卷
func AddTestPaper(testpaper *TestPaper) {
	DB.Create(&testpaper)
}

// FindChoiceTP 透過 ID 取得 Choices
func FindTestPaper(id uint) (TestPaper, error) {
	var tp TestPaper
	if err := DB.Where("id = ?", id).First(&tp).Error; err != nil {
		return TestPaper{}, err
	}
	return tp, nil
}

// EditTP 修改
func EditTP(tp TestPaper) {
	DB.Where("id = ?", tp.ID).Save(&tp)
}

// DeleteTP 刪除
func DeleteTP(tp TestPaper) {
	DB.Where("id = ?", tp.ID).Delete(&tp)
}

// DeleteChoiceTP 刪除
func DeleteChoiceTP(choiceTP ChoiceTP) {
	DB.Where("id = ?", choiceTP.ID).Delete(&choiceTP)
}

// DeleteClozeTP 刪除
func DeleteClozeTP(clozeTP ClozeTP) {
	DB.Where("id = ?", clozeTP.ID).Delete(&clozeTP)
}
