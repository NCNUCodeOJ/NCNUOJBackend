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

		nor.GET("/find/byprobid", view.GetProblemByProblemId)     // 用problemid 查詢 problem
		nor.GET("/find/bytagid", view.GetProblemByTagId)          // 查這個標籤，有哪些problem有用
		nor.GET("/find/byprobname", view.GetProblemByProblemName) // 用題目名稱查詢
		nor.POST("/createproblem", view.CreateProblem)            // 創建題目
		nor.POST("/deleteproblem", view.DeleteProblemById)        // 用problemid 刪掉 problem
		nor.PUT("/editproblem", view.EditProblemByProblemId)      // 編輯更新problem
	}

	return r
}
