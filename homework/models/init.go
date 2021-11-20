package models

import (
	//"fmt"

	"context"
	"log"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Ctx = context.Background()
var RC *redis.Client

//DB 資料庫連接
var DB *gorm.DB

const (
	RedisToken        = ""
	RedisProblems     = "redis_problems"
	RedisSamples      = "redis_samples"
	RedisTags         = "redis_tags"
	RedisTag2Problems = "redis_tag_2_problems"
)

//Setup 資料庫連接設定
func Setup() {
	var err error
	DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	AutoMigrateAll()

}

//AutoMigrateAll 自動產生 table
func AutoMigrateAll() {
	DB.AutoMigrate(&Problem{})
	DB.AutoMigrate(&Tag{})
	DB.AutoMigrate(&Sample{})
	DB.AutoMigrate(&Tag2Problem{})

}
func SetRc() {
	RC = newClient()

	// 初始化清空所有Redis

	_, err := RC.Del(Ctx, RedisProblems, RedisSamples, RedisTag2Problems, RedisTags).Result()
	if err != nil {
		panic(err)
	}

}
func newClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       1,
	})

	pong, err := client.Ping(Ctx).Result()
	log.Println(pong)
	if err != nil {
		log.Fatalln(err)
	}
	return client
}
