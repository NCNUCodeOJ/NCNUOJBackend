package router

import (
	"NCNUOJBackend/testPaper/view"

	"github.com/gin-gonic/gin"
)

// /testpaper/{id}/topic/{id}/question/{sort}}

// SetupRouter index
// *gin.Engine 是要回傳的型態
func SetupRouter() *gin.Engine {
	baseURL := "api/v1"
	r := gin.Default()
	// testpaper 測驗卷
	// Group 可以讓網址延伸，不用重複寫
	testpaper := r.Group(baseURL + "/testpaper")
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
		// /testpaper/{testpaperid}/topic/{topicid}/question/{topic裡的第幾題}
		// 替換大題中的題目
		// topic.PATCH("/:topicID/question/topicid", view.EditTopic)
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
	// answer 選項答案
	answer := r.Group(baseURL + "/question/:questionID/answer")
	{
		answer.POST("", view.AddAnswer)
		answer.GET("", view.GetAllAnswers)
		answer.GET("/:answerID", view.GetAnswer)
		answer.PATCH("/:answerID", view.EditAnswer)
		answer.DELETE("/:answerID", view.DeleteAnswer)
		// /testpaper/{testpaperid}/answer/{answerid}/question/{answer裡的第幾題}
		// 替換大題中的題目
		// answer.PATCH("/:answerID/question/answerid", view.Editanswer)
	}
	return r
}
