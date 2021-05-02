package view

import (
	"NCNUOJBackend/testPaper/models"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vincentinttsh/replace"
	"github.com/vincentinttsh/zero"
)

// SetCloze 新增填充題
func SetCloze(c *gin.Context) {
	// 使用者傳過來的檔案格式(選擇題題目、出題者、範圍、出處)
	var data struct {
		Question *string `json:"question"`
		AuthorID *uint   `json:"authorID"`
		Layer    *uint   `json:"layer"`
		Source   *uint   `json:"source"`
	}
	var cloze models.Cloze
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "選擇題未按照格式填寫",
		})
		return
	}
	// 如果有空值，則回傳 false
	if zero.IsZero(data) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "選擇題所有欄位不可為空值",
		})
		return
	}
	cloze.Question = *data.Question
	cloze.AuthorID = *data.AuthorID
	cloze.Layer = *data.Layer
	cloze.Source = *data.Source
	models.AddCloze(&cloze)
	c.JSON(http.StatusOK, gin.H{
		"message": "新增選擇題題目成功",
	})
}

// SetClozeanswer 新增填充題答案
func SetClozeAnswer(c *gin.Context) {
	// 使用者傳過來的檔案格式(選項內容、是否為正確答案、對應的題目)
	var data struct {
		Content *string `json:"content"`
		ClozeID *uint   `json:"clozeID"`
	}
	var answer models.ClozeAnswer
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "選擇題選項未按照格式填寫",
		})
		return
	}
	// 如果有空值，則回傳 false
	if zero.IsZero(data) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "選擇題選項所有欄位不可為空值",
		})
		return
	}
	answer.Content = *data.Content
	answer.ClozeID = *data.ClozeID
	models.AddClozeAnswer(&answer)
	c.JSON(http.StatusOK, gin.H{
		"message": "新增選擇題選項成功",
	})
}

// FindCloze 查詢填充題
func FindCloze(c *gin.Context) {
	data := struct {
		ID uint `json:"ID"`
	}{}
	c.Bind(&data)
	if cloze, err := models.FindCloze(data.ID); err == nil && cloze.ID != 0 {
		c.JSON(http.StatusOK, gin.H{
			"id":       cloze.ID,
			"question": cloze.Question,
			"authorID": cloze.AuthorID,
			"layer":    cloze.Layer,
			"source":   cloze.Source,
		})
		return
	}
}

// Edit 修改選擇題題目
func EditCloze(c *gin.Context) {
	data := struct {
		ID       *uint   `json:"ID"`
		Question *string `json:"question"`
		AuthorID *uint   `json:"authorID"`
		Layer    *uint   `json:"layer"`
		Source   *uint   `json:"source"`
	}{}
	c.BindJSON(&data)
	if cloze, err := models.FindCloze(*data.ID); err == nil && cloze.ID != 0 {
		replace.Replace(&cloze, &data)
		c.JSON(http.StatusOK, gin.H{
			"message": "修改成功",
		})
		return
	}
}

// FindCloze 查詢選擇題
func FindAnswer(c *gin.Context) {
	data := struct {
		ID uint `json:"ID"`
	}{}
	c.Bind(&data)
	if ans, err := models.FindAnswer(data.ID); err == nil && ans.ID != 0 {
		c.JSON(http.StatusOK, gin.H{
			"id":      ans.ID,
			"content": ans.Content,
			"clozeID": ans.ClozeID,
		})
		return
	}
}

// Edit 修改選擇題答案
func EditAnswer(c *gin.Context) {
	data := struct {
		ID      *uint   `json:"ID"`
		Content *string `json:"content"`
		ClozeID *uint   `json:"clozeID"`
	}{}
	c.BindJSON(&data)
	if ans, err := models.FindAnswer(*data.ID); err == nil && ans.ID != 0 {
		replace.Replace(&ans, &data)
		c.JSON(http.StatusOK, gin.H{
			"message": "修改成功",
		})
		return
	}
}

// DeleteCloze 刪除填充題
func DeleteCloze(c *gin.Context) {
	data := struct {
		ID uint `json:"ID"`
	}{}
	c.BindJSON(&data)
	if cloze, err := models.FindCloze(data.ID); err == nil && cloze.ID != 0 {
		models.DeleteCloze(cloze)
		c.JSON(http.StatusOK, gin.H{
			"message": "填充題刪除成功",
		})
		return
	}
}

// DeleteAnswer 刪除答案
func DeleteAnswer(c *gin.Context) {
	data := struct {
		ID uint `json:"ID"`
	}{}
	c.BindJSON(&data)
	if ans, err := models.FindAnswer(data.ID); err == nil && ans.ID != 0 {
		models.DeleteAnswer(ans)
		c.JSON(http.StatusOK, gin.H{
			"message": "答案刪除成功",
		})
		return
	}
}
