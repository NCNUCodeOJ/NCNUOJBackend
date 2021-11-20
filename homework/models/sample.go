package models

import (
	"gorm.io/gorm"
)

//Sample Database - database
type Sample struct {
	gorm.Model
	Input     string `gorm:"type:text;"`
	Output    string `gorm:"type:text;"`
	ProblemId uint   `gorm:"NOT NULL;"`
}

//AddSample 增加範例
func AddSample(sample *Sample) {
	DB.Create(&sample)
}

//DeleteSample 用problem id 直接有這個 problemid 的範例
func DeleteSample(id uint) (err error) {
	err = DB.Where("problem_id = ?", id).Delete(&Sample{}).Error
	return
}

//EditForDeleteSample 編輯，刪除更新後不存在的舊資料
func EditForDeleteSample(sample_id, problem_id uint) {
	DB.Where("id = ? AND problem_id = ?", sample_id, problem_id).Delete(&Sample{})
}

//SampleDetailByProblemId 用problemid找sample
func SampleDetailByProblemId(id uint) ([]Sample, error) {
	var sample []Sample
	if err := DB.Where("problem_id = ?", id).Find(&sample).Error; err != nil {
		return []Sample{}, err
	}

	return sample, nil
}
