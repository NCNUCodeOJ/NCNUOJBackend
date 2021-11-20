package models

import (
	"gorm.io/gorm"
)

// Tag Database - database
type Tag2Problem struct {
	gorm.Model
	TagId     uint `gorm:"NOT NULL"`
	ProblemId uint `gorm:"NOT NULL"`
}

//AddTag2Table 創建 tag to problem
func AddTag2Problem(tag2table *Tag2Problem) {
	DB.Create(&tag2table)
}

//DeleteByProblemId 用 problem id 刪除 所有與這個problem id 有關的 row
func DeleteTag2ProblemByProb(id uint) (err error) {
	err = DB.Where("problem_id = ?", id).Delete(&Tag2Problem{}).Error
	return
}

//DeleteTag2TableByTag  用 tag id 刪除 所有與這個 tag id 有關的 row
func DeleteTag2ProblemByTag(id uint) {
	DB.Where("tag_id = ?", id).Delete(&Tag2Problem{})
}

//EditForDeleteTag2Table 編輯，刪除更新後不存在舊資料
func EditForDeleteTag2Table(tag_id, problem_id uint) {
	DB.Where("tag_id = ? AND problem_id = ?", tag_id, problem_id).Delete(&Tag2Problem{})
}

//TagDetailByProblemId 用　problemid　找有哪些 tag 在這個 problem 上
func TagDetailByProblemId(id uint) ([]Tag2Problem, error) {
	var tag2problem []Tag2Problem
	if err := DB.Where("problem_id = ?", id).Find(&tag2problem).Error; err != nil {
		return []Tag2Problem{}, err
	}

	return tag2problem, nil

}

//ProblemDetailByTagId 用　tagid　找有哪些 problem 有用到這個 tag
func ProblemDetailByTagId(id uint) ([]Tag2Problem, error) {
	var tag2problem []Tag2Problem
	if err := DB.Where("tag_id = ?", id).Find(&tag2problem).Error; err != nil {
		return []Tag2Problem{}, err
	}
	return tag2problem, nil

}
