package models

import "gorm.io/gorm"

//User Database - database
type User struct {
	gorm.Model
	SchoolID  int    `gorm:"NOT NULL;"`
	StudentID string `gorm:"type:varchar(15) NOT NULL;"`
	Email     string `gorm:"type:varchar(40) NOT NULL;"`
	UserName  string `gorm:"type:varchar(20) NOT NULL;"`
	Password  string `gorm:"type:varchar(100) NOT NULL;"`
	RealName  string `gorm:"type:varchar(30)"`
	Admin     bool   `gorm:"default:false"`
}

// UserDetailByID 透過 id 取得 username
func UserDetailByID(id uint) (user User) {
	DB.Where("id = ?", id).First(&user)
	return
}

// UserDetailByUserName 透過 UserName 取得 username
func UserDetailByUserName(name string) (User, error) {
	var user User
	if err := DB.Where("user_name = ?", name).First(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

// AddUser 新增 user
func AddUser(user *User) {
	DB.Create(&user)
}
