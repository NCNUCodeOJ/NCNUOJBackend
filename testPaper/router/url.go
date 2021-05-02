package router

import (
	"NCNUOJBackend/testPaper/view"

	"github.com/gin-gonic/gin"
)

// SetupRouter index
// *gin.Engine 是要回傳的型態
func SetupRouter() *gin.Engine {
	baseURL := "api/v1"
	r := gin.Default()
	// Group 可以讓網址延伸，不用重複寫
	nor := r.Group(baseURL + "/testpaper")
	{
		// 完整的網址是 http://localhost:8080/api/v1/testpaper/admin
		// testpaper 測驗卷
		nor.POST("/admin", view.SetTestPaper)
		nor.POST("/admin/choicetp", view.ChoiceTP)
		nor.POST("/admin/clozetp", view.ClozeTP)
		nor.GET("/find", view.FindTestPaper)
		nor.POST("/edit", view.EditTP)
		nor.DELETE("/delete", view.DeleteTP)
		nor.DELETE("/delete/choicetp", view.DeleteChoiceTP)
		nor.DELETE("/delete/clozetp", view.DeleteClozeTP)
		// Choice 選擇題
		nor.POST("/setchoice", view.SetChoice)
		nor.POST("/setchoice/option", view.SetChoiceOption)
		nor.GET("/find/choice", view.FindChoice)
		nor.GET("/find/choice/option", view.FindOption)
		nor.POST("/edit/choice", view.EditChoice)
		nor.POST("/edit/choice/option", view.EditOption)
		nor.DELETE("/delete/choice", view.DeleteChoice)
		nor.DELETE("/delete/choice/option", view.DeleteOption)
		// Cloze 填充題
		nor.POST("/setcloze", view.SetCloze)
		nor.POST("/setcloze/answer", view.SetClozeAnswer)
		nor.GET("/find/cloze", view.FindCloze)
		nor.GET("/find/cloze/answer", view.FindAnswer)
		nor.POST("/edit/cloze", view.EditCloze)
		nor.POST("/edit/cloze/answer", view.EditAnswer)
		nor.DELETE("/delete/cloze", view.DeleteCloze)
		nor.DELETE("/delete/cloze/answer", view.DeleteAnswer)
	}
	return r
}
