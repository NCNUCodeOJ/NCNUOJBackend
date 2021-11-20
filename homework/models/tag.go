package models

import "gorm.io/gorm"

// Tag Database - database
type Tag struct {
	gorm.Model
	Name string `gorm:"type:varchar(20) NOT NULL;"`
}

//AddTag 創建 tag
func AddTag(tag *Tag) {
	DB.Create(&tag)
}

//CheckTag 檢查表格裡面有沒有重複的 tag
func CheckTag(name string) (Tag, error) {
	var tag Tag
	if err := DB.Where("name = ?", name).First(&tag).Error; err != nil {
		return Tag{}, err
	}
	return tag, nil

}

//TagDetailByTagId 用 tagid 查 tag
func TagDetailByTagId(id uint) (Tag, error) {
	var tag Tag
	if err := DB.Where("id = ?", id).Find(&tag).Error; err != nil {
		return Tag{}, err
	}
	return tag, nil
}
