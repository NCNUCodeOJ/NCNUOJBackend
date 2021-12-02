package models

import (
	"gorm.io/gorm"
)

// 第幾大題(要放在測驗卷的)
type Topic struct {
	gorm.Model
	Description string `gorm:"type:text;"`
	TestPaperID uint   `gorm:"NOT NULL;"`
	Sort        uint   `gorm:"NOT NULL;"`
	// 大題敘述
	// 對應的測驗卷
	// 排序(這是第幾大題)
}

// AddTopic 新增大題
func AddTopic(topic *Topic) {
	DB.Create(&topic)
}

// GetAllTopics 取得所有 topic
func GetAllTopics() (topics []Topic, err error) {
	err = DB.Find(&topics).Error
	return
}

// GetTopic 透過 testPaperID 取得 Topic
func GetTopic(testpaperID uint) (Topic, error) {
	var topic Topic
	if err := DB.Where("id = ?", testpaperID).First(&topic).Error; err != nil {
		return Topic{}, err
	}
	return topic, nil
}

// EditTopic 修改
func EditTopic(topic *Topic) (err error) {
	DB.Where("id = ?", topic.TestPaperID).Save(&topic)
	return
}

// DeleteTopic 刪除
func DeleteTopic(topic Topic) (err error) {
	DB.Where("id = ?", topic.TestPaperID).Delete(&topic)
	return
}
