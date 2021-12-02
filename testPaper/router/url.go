package router

import (
	"NCNUOJBackend/testPaper/view"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// /testpaper/{id}/topic/{id}/question/{sort}}

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
	r := gin.Default()
	// testpaper 測驗卷
	// Group 可以讓網址延伸，不用重複寫
	testpaper := r.Group(baseURL + "/testpaper")
	testpaper.Use(authMiddleware.MiddlewareFunc())
	testpaper.Use(getUserID())
	{
		testpaper.POST("", view.AddTestPaper)
		testpaper.GET("", view.GetAllTestPapers)
		testpaper.GET("/:testpaperID", view.GetTestPaper)
		testpaper.PATCH("/:testpaperID", view.EditTestPaper)
		testpaper.DELETE("/:testpaperID", view.DeleteTestPaper)
	}
	// topic 大題
	topic := r.Group(baseURL + "/testpaper/:testpaperID/topic")
	{
		topic.POST("", view.AddTopic)
		topic.GET("", view.GetAllTopics)
		topic.GET("/:topicID", view.GetTopic)
		topic.PATCH("/:topicID", view.EditTopic)
		topic.DELETE("/:topicID", view.DeleteTopic)
	}
	// Question 題目
	question := r.Group(baseURL + "/question")
	{
		question.POST("", view.AddQuestion)
		// question.GET("", view.GetAllQuestions)
		question.GET("/:questionID", view.GetQuestion)
		// question.GET("/:id/:sort", view.GetQuestion)
		question.PATCH("/:questionID", view.EditQuestion)
		question.DELETE("/:questionID", view.DeleteQuestion)
		// 對使用者而言，一個問題就是一個物件
	}
	// answer 選項/答案
	answer := r.Group(baseURL + "/question/:questionID/answer")
	{
		answer.POST("", view.AddAnswer)
		answer.GET("", view.GetAllAnswers)
		answer.GET("/:answerID", view.GetAnswer)
		answer.PATCH("/:answerID", view.EditAnswer)
		answer.DELETE("/:answerID", view.DeleteAnswer)
	}
	return r
}

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
