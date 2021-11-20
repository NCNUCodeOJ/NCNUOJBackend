package models

import "gorm.io/gorm"

//Homework Database - database
type Homework struct {
	gorm.Model
	HwName    string `gorm:"type:text;"`
	ProblemID []uint
}
