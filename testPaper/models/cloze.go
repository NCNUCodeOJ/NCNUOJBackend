package models

import "gorm.io/gorm"

//User Database - database
type Cloze struct {
	gorm.Model
	Question string `gorm:"NOT NULL;"`
	AuthorID uint
	Layer    uint
	Source   uint
	// 填充題題目
	// 出題者
	// 層級(校內、區域、全國)
	// 出處(學校ID、單位ID)
}

// 填充題答案
type ClozeAnswer struct {
	gorm.Model
	Content string `gorm:"NOT NULL;"`
	ClozeID uint
	// 正確答案
	// 對應的填充題題目
}

// 填充題們(要放在測驗卷的)
type ClozeTP struct {
	gorm.Model
	ClassID uint
	ClozeID uint
	Sort    uint
	// 對應的課堂
	// 對應的填充題題目
	// 排序
}

// AddCloze 新增填充題題目
func AddCloze(cloze *Cloze) {
	DB.Create(&cloze)
}

// AddClozeAnswer 新增填充題答案
func AddClozeAnswer(answer *ClozeAnswer) {
	DB.Create(&answer)
}

// AddClozes 新增填充題們
func AddClozes(clozeTP *ClozeTP) {
	DB.Create(&clozeTP)
}

// FindCloze 透過 ID 取得 Cloze
func FindCloze(id uint) (Cloze, error) {
	var cloze Cloze
	if err := DB.Where("id = ?", id).First(&cloze).Error; err != nil {
		return Cloze{}, err
	}
	return cloze, nil
}

// FindOption 透過 ID 取得 ClozeOption
func FindAnswer(id uint) (ClozeAnswer, error) {
	var clozeanswer ClozeAnswer
	if err := DB.Where("id = ?", id).First(&clozeanswer).Error; err != nil {
		return ClozeAnswer{}, err
	}
	return clozeanswer, nil
}

// FindClozeTP 透過 ID 取得 ClozeTP
func FindClozeTP(id uint) (ClozeTP, error) {
	var clozeTP ClozeTP
	if err := DB.Where("id = ?", id).First(&clozeTP).Error; err != nil {
		return ClozeTP{}, err
	}
	return clozeTP, nil
}

// EditCloze 修改
func EditCloze(cloze Cloze) {
	DB.Where("id = ?", cloze.ID).Save(&cloze)
}

// EditAns 修改
func EditAns(ans ClozeAnswer) {
	DB.Where("id = ?", ans.ID).Save(&ans)
}

// DeleteCloze 刪除
func DeleteCloze(cloze Cloze) {
	DB.Where("id = ?", cloze.ID).Delete(&cloze)
}

// DeleteAnswer 刪除
func DeleteAnswer(ans ClozeAnswer) {
	DB.Where("id = ?", ans.ID).Delete(&ans)
}
