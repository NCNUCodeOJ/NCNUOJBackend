package router

import (
	"NCNUOJBackend/homework/view"

	"github.com/gin-gonic/gin"
)

// SetupRouter index
func SetupRouter() *gin.Engine {
	baseURL := "api/v1"
	r := gin.Default()
	nor := r.Group(baseURL + "/prob")
	{

		nor.GET("/problem/:probid", view.GetProblemByProblemId) // 用problemid 查詢 problem
		nor.GET("/tag/:tagid", view.GetProblemByTagId)          // 查這個標籤，有哪些problem有用
		nor.GET("/list_of_problem", view.GetProblemList)        //列出所有題目id
		nor.POST("/createproblem", view.CreateProblem)          // 創建題目
		nor.DELETE("/delete/:probid", view.DeleteProblemById)   // 用problemid 刪掉 problem
		nor.PUT("/edit/:probid", view.EditProblemByProblemId)   // 編輯更新problem
	}

	return r
}
