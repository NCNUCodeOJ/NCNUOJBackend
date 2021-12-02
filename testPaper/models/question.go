package models

import "gorm.io/gorm"

// 選擇題type:text;
type Question struct {
	gorm.Model
	Question   string `gorm:"type:text;"`
	AuthorID   uint   `gorm:"NOT NULL;"`
	Layer      uint   `gorm:"NOT NULL;"`
	Source     uint   `gorm:"NOT NULL;"`
	Difficulty uint   `gorm:"NOT NULL;"`
	Type       uint   `gorm:"NOT NULL;"`
	// 選擇題題目
	// 出題者
	// 層級(校內、區域、全國)
	// 題目出處(學校 ID、單位 ID)
	// 難易度
	// 類型(多選、單選、填充)
	// 答案/選項
}

// Answer
type Answer struct {
	gorm.Model
	Content    string `gorm:"type:text;"`
	Correct    bool   `gorm:"NOT NULL;"`
	QuestionID uint   `gorm:"NOT NULL;"`
	Sort       uint   `gorm:"NOT NULL;"`
	// 內容
	// 是否為正確答案
	// 對應的題目
	// 這是第幾個選項(若為填充題則填 -1)
}

// AddQuestion 新增選擇題題目
func AddQuestion(question *Question) {
	DB.Create(&question)
}

// AddAnswer 新增選項
func AddAnswer(answer *Answer) {
	DB.Create(&answer)
}

// GetQuestionByID 透過 ID 取得 Question
func GetQuestion(id uint) (Question, error) {
	var question Question
	if err := DB.Where("id = ?", id).First(&question).Error; err != nil {
		return Question{}, err
	}
	return question, nil
}

// GetAnswer 透過 questionID 取得 Answer
func GetAnswer(questionID uint) (Answer, error) {
	var answer Answer
	if err := DB.Where("id = ?", questionID).First(&answer).Error; err != nil {
		return Answer{}, err
	}
	return answer, nil
}

// GetAllAnswers 取得所有 answer
func GetAllAnswers() (answers []Answer, err error) {
	err = DB.Find(&answers).Error
	return
}

// EditQuestion 修改題目
func EditQuestion(question *Question) (err error) {
	err = DB.Where("id = ?", question.ID).Save(&question).Error
	return
}

// EditAnswer 修改選項
func EditAnswer(answer *Answer) (err error) {
	DB.Where("id = ?", answer.QuestionID).Save(&answer)
	return
}

// DeleteQuestion 刪除題目
func DeleteQuestion(question Question) (err error) {
	err = DB.Where("id = ?", question.ID).Delete(&question).Error
	return
}

// DeleteAnswer 刪除選項
func DeleteAnswer(answer Answer) (err error) {
	// err = DB.Where("questionID = ?", answer.QuestionID).Where("sort = ?", answer.Sort).Delete(&answer).Error
	err = DB.Where("id = ?", answer.QuestionID).Delete(&answer).Error
	return
}
