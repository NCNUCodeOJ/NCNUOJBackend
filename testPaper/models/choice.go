package models

import "gorm.io/gorm"

// 選擇題
type Choice struct {
	gorm.Model
	Question string `gorm:"NOT NULL;"`
	AuthorID uint
	Layer    uint
	Source   uint
	// 選擇題題目
	// 出題者
	// 層級(校內、區域、全國)
	// 題目出處(學校ID、單位ID)
}

// 選項
type ChoiceOption struct {
	gorm.Model
	Content  string `gorm:"NOT NULL;"`
	Answer   bool
	ChoiceID uint
	// 選項內容
	// 是否為正確答案
	// 對應的選擇題題目
}

// 選擇題們(要放在測驗卷的)
type ChoiceTP struct {
	gorm.Model
	ClassID  uint
	ChoiceID uint
	Sort     uint
	// 對應的課堂
	// 對應的選擇題題目
	// 排序
}

// AddChoice 新增選擇題題目
func AddChoice(choice *Choice) {
	DB.Create(&choice)
}

// AddChoiceOption 新增選擇題選項
func AddChoiceOption(option *ChoiceOption) {
	DB.Create(&option)
}

// AddChoices 新增選擇題們
func AddChoiceTP(choices *ChoiceTP) {
	DB.Create(&choices)
}

// FindChoiceByID 透過 ID 取得 Choice
func FindChoice(id uint) (Choice, error) {
	var choice Choice
	if err := DB.Where("id = ?", id).First(&choice).Error; err != nil {
		return Choice{}, err
	}
	return choice, nil
}

// FindOptionByID 透過 ID 取得 ChoiceOption
func FindOption(id uint) (ChoiceOption, error) {
	var option ChoiceOption
	if err := DB.Where("id = ?", id).First(&option).Error; err != nil {
		return ChoiceOption{}, err
	}
	return option, nil
}

// FindChoiceTP 透過 ID 取得 Choices
func FindChoiceTP(id uint) (ChoiceTP, error) {
	var choiceTP ChoiceTP
	if err := DB.Where("id = ?", id).First(&choiceTP).Error; err != nil {
		return ChoiceTP{}, err
	}
	return choiceTP, nil
}

// EditChoice 修改
func EditChoice(choice Choice) {
	DB.Where("id = ?", choice.ID).Save(&choice)
}

// EditChoice 修改
func EditOption(opt ChoiceOption) {
	DB.Where("id = ?", opt.ID).Save(&opt)
}

// DeleteChoice 刪除
func DeleteChoice(choice Choice) {
	DB.Where("id = ?", choice.ID).Delete(&choice)
}

// DeleteOption 刪除
func DeleteOption(opt ChoiceOption) {
	DB.Where("id = ?", opt.ID).Delete(&opt)
}
