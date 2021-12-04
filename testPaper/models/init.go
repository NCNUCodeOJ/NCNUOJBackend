package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//DB 資料庫連接
var DB *gorm.DB

//Setup 資料庫連接設定
func Setup() {
	var err error
	DB, err = gorm.Open(sqlite.Open("testpaper.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	AutoMigrateAll()
}

//AutoMigrateAll 自動產生 table
func AutoMigrateAll() {
	DB.AutoMigrate(&TestPaper{})
	DB.AutoMigrate(&QuestionTopic{})
	DB.AutoMigrate(&Topic{})
	DB.AutoMigrate(&Question{})
	DB.AutoMigrate(&Option{})
}

// var Ctx = context.Background()
// var RC *redis.Client

// func Client() {
// 	rdb := redis.NewClient(&redis.Options{
// 		Addr:     "localhost:6379",
// 		Password: "", // no password set
// 		DB:       0,  // use default DB
// 	})
// 	pong, err := rdb.Ping(Ctx).Result()
// 	if err != nil {
// 		fmt.Println("Redis connection fail：", pong, err)
// 		return
// 	}
// 	fmt.Println("Redis connection successfully：", pong)
// }
