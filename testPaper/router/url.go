package router

import (
	"NCNUOJBackend/testPaper/view"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// var token = "Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6ZmFsc2UsImV4cCI6NDc5MTA4MjEyMywiaWQiOiI3MTI0MTMxNTQxOTcxMTA3ODYiLCJvcmlnX2lhdCI6MTYzNzQ4MjEyMywidXNlcm5hbWUiOiJ0ZXN0X3VzZXIifQ.pznOSok8X7qv6FSIihJnma_zEy70TerzOs0QDZOq_4RPYOKSEOOYTZ9-VLm2P9XRldS17-7QrLFwjjfXyCodtA"

func getUserID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(jwt.ExtractClaims(c)["id"].(string))
		if err != nil {
			c.Abort()
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "系統錯誤",
				"error":   err.Error(),
			})
		} else {
			c.Set("userID", uint(id))
			c.Next()
		}
	}
}

// SetupRouter index
// *gin.Engine 是要回傳的型態
func SetupRouter() *gin.Engine {
	if gin.Mode() == "test" {
		err := godotenv.Load(".env.test")
		if err != nil {
			log.Println("Error loading .env file")
		}
	} else if gin.Mode() == "debug" {
		err := godotenv.Load()
		if err != nil {
			log.Println("Error loading .env file")
		}
	}
	fmt.Println(os.Getenv("SECRET_KEY"))
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:            "NCNUOJ",
		SigningAlgorithm: "HS512",
		Key:              []byte(os.Getenv("SECRET_KEY")),
		MaxRefresh:       time.Hour,
		TimeFunc:         time.Now,
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	baseURL := "api/v1"
	privateURL := "api/private/v1"
	r := gin.Default()
	// testpaper 測驗卷
	// Group 可以讓網址延伸，不用重複寫
	testpaper := r.Group(privateURL + "/testpaper")
	testpaper.Use(authMiddleware.MiddlewareFunc())
	testpaper.Use(getUserID())
	{
		testpaper.POST("", view.CreateTestPaper)
		testpaper.GET("", view.ListTestPapers)
		testpaper.GET("/:testpaperID", view.GetTestPaperByID)
		testpaper.PATCH("/:testpaperID", view.UpdateTestPaper)
		testpaper.DELETE("/:testpaperID", view.DeleteTestPaper)
	}
	// topic 大題
	topic := r.Group(privateURL + "/testpaper/:testpaperID/topic")
	{
		topic.POST("", view.CreateTopic)
		topic.GET("", view.ListTopics)
		topic.GET("/:sort", view.GetTopicBySort)
		topic.PATCH("/:sort", view.UpdateTopic)
		topic.DELETE("/:sort", view.DeleteTopic)
	}
	// Question 題目 (含選項/答案)
	question := r.Group(baseURL + "/question")
	question.Use(authMiddleware.MiddlewareFunc())
	question.Use(getUserID())
	{
		question.POST("", view.CreateQuestion)
		question.GET("", view.ListQuestions)
		question.GET("/:questionID", view.GetQuestion)
		question.PATCH("/:questionID", view.UpdateQuestion)
		question.DELETE("/:questionID", view.DeleteQuestion)
		// 對使用者而言，一個問題就是一個物件
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Page not found"})
	})
	return r
}
