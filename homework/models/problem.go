package models

import "gorm.io/gorm"

//Problem Database - database
type Problem struct {
	gorm.Model
	ProblemName      string `gorm:"type:text;"`
	Description      string `gorm:"type:text;"`
	InputDescription string `gorm:"type:text;"`
	OutputDescripton string `gorm:"type:text"`
	Author           uint   `gorm:"NOT NULL;"`
	MemoryLimit      uint   `gorm:"NOT NULL;"`
	Cputime          uint   `gorm:"NOT NULL;"`
	Layer            uint8  `gorm:"NOT NULL;"`
}

//AddProblem 創建題目
func AddProblem(problem *Problem) (err error) {
	err = DB.Create(&problem).Error
	return
}

//UpdateProblem 更新題目
func UpdateProblem(problem *Problem) (err error) {
	err = DB.Where("id = ?", problem.ID).Save(&problem).Error
	return
}

//DeleteProblem 刪除題目
func DeleteProblem(id uint) (err error) {
	err = DB.Delete(&Problem{}, id).Error
	return
}

//ProblemDetailByProblemId 查詢題目用problemid
func ProblemDetailByProblemId(id uint) (Problem, error) {
	var problem Problem
	if err := DB.Where("id = ?", id).First(&problem).Error; err != nil {
		return Problem{}, err
	}
	return problem, nil
}

//ProblemDetailByProblemId 查詢題目用problem name
func ProblemDetailByProblemName(name string) (Problem, error) {
	var problem Problem
	if err := DB.Where("problem_name = ?", name).First(&problem).Error; err != nil {
		return Problem{}, err
	}
	return problem, nil
}
