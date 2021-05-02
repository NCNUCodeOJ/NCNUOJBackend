package models

import "gorm.io/gorm"

// Tag Database - database
type Tag2Table struct {
	gorm.Model
	TagId     uint `gorm:"NOT NULL"`
	ProblemId uint `gorm:"NOT NULL"`
}

//AddTag2Table 創建 tag to problem
func AddTag2Table(tag2table *Tag2Table) {
	DB.Create(&tag2table)
}

//DeleteByProblemId 用 problem id 刪除 所有與這個problem id 有關的 row
func DeleteTag2TableByProb(id uint) (err error) {
	err = DB.Where("problem_id = ?", id).Delete(&Tag2Table{}).Error
	return
}

//DeleteTag2TableByTag  用 tag id 刪除 所有與這個 tag id 有關的 row
func DeleteTag2TableByTag(id uint) {
	DB.Where("tag_id = ?", id).Delete(&Tag2Table{})
}

//EditForDeleteTag2Table 編輯，刪除更新後不存在舊資料
func EditForDeleteTag2Table(tag_id, problem_id uint) {
	DB.Where("tag_id = ? AND problem_id = ?", tag_id, problem_id).Delete(&Tag2Table{})
}

//TagDetailByProblemId 用　problemid　找有哪些 tag 在這個 problem 上
func TagDetailByProblemId(id uint) ([]Tag2Table, error) {
	var tag2table []Tag2Table
	if err := DB.Where("problem_id = ?", id).Find(&tag2table).Error; err != nil {
		return []Tag2Table{}, err
	}
	return tag2table, nil

}

//ProblemDetailByTagId 用　tagid　找有哪些 problem 有用到這個 tag
func ProblemDetailByTagId(id uint) ([]Tag2Table, error) {
	var tag2table []Tag2Table
	if err := DB.Where("tag_id = ?", id).Find(&tag2table).Error; err != nil {
		return []Tag2Table{}, err
	}
	return tag2table, nil

}
