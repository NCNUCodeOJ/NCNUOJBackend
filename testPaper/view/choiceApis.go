package view

import (
	"NCNUOJBackend/testPaper/models"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vincentinttsh/replace"
	"github.com/vincentinttsh/zero"
)

// SetChoice 新增選擇題
func SetChoice(c *gin.Context) {
	// 使用者傳過來的檔案格式(選擇題題目、出題者、範圍、出處)
	var data struct {
		Question *string `json:"question"`
		AuthorID *uint   `json:"authorID"`
		Layer    *uint   `json:"layer"`
		Source   *uint   `json:"source"`
	}
	var choice models.Choice
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
	choice.Question = *data.Question
	choice.AuthorID = *data.AuthorID
	choice.Layer = *data.Layer
	choice.Source = *data.Source
	models.AddChoice(&choice)
	c.JSON(http.StatusOK, gin.H{
		"message": "新增選擇題題目成功",
	})
}

// SetChoiceOption 新增選擇題選項
func SetChoiceOption(c *gin.Context) {
	// 使用者傳過來的檔案格式(選項內容、是否為正確答案、對應的題目)
	var data struct {
		Content  *string `json:"content"`
		Answer   *bool   `json:"answer"`
		ChoiceID *uint   `json:"choiceID"`
	}
	var option models.ChoiceOption
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
	option.Content = *data.Content
	option.Answer = *data.Answer
	option.ChoiceID = *data.ChoiceID
	models.AddChoiceOption(&option)
	c.JSON(http.StatusOK, gin.H{
		"message": "新增選擇題選項成功",
	})
}

// FindChoice 查詢選擇題
func FindChoice(c *gin.Context) {
	data := struct {
		ID uint `json:"ID"`
	}{}
	c.Bind(&data)
	if choice, err := models.FindChoice(data.ID); err == nil && choice.ID != 0 {
		c.JSON(http.StatusOK, gin.H{
			"id":       choice.ID,
			"question": choice.Question,
			"authorID": choice.AuthorID,
			"layer":    choice.Layer,
			"source":   choice.Source,
		})
		return
	}
}

// Edit 修改
func EditChoice(c *gin.Context) {
	data := struct {
		ID       *uint   `json:"ID"`
		Question *string `json:"question"`
		AuthorID *uint   `json:"authorID"`
		Layer    *uint   `json:"layer"`
		Source   *uint   `json:"source"`
	}{}
	c.BindJSON(&data)
	if choice, err := models.FindChoice(*data.ID); err == nil && choice.ID != 0 {
		replace.Replace(&choice, &data)
		models.EditChoice(choice)
		c.JSON(http.StatusOK, gin.H{
			"message": "修改成功",
		})
		return
	}
}

// FindOption 查詢選擇題選項
func FindOption(c *gin.Context) {
	data := struct {
		ID uint `json:"ID"`
	}{}
	c.Bind(&data)
	if opt, err := models.FindOption(data.ID); err == nil && opt.ID != 0 {
		c.JSON(http.StatusOK, gin.H{
			"id":       opt.ID,
			"Content":  opt.Content,
			"Answer":   opt.Answer,
			"ChoiceID": opt.ChoiceID,
		})
		return
	}
}

// Edit 修改
func EditOption(c *gin.Context) {
	data := struct {
		ID       *uint   `json:"ID"`
		Content  *string `json:"content"`
		Answer   *bool   `json:"answer"`
		ChoiceID *uint   `json:"choiceID"`
	}{}
	c.BindJSON(&data)
	if opt, err := models.FindOption(*data.ID); err == nil && opt.ID != 0 {
		replace.Replace(&opt, &data)
		models.EditOption(opt)
		c.JSON(http.StatusOK, gin.H{
			"message": "修改成功",
		})
		return
	}
}

// DeleteChoice 刪除選擇題
func DeleteChoice(c *gin.Context) {
	data := struct {
		ID uint `json:"ID"`
	}{}
	c.BindJSON(&data)
	if choice, err := models.FindChoice(data.ID); err == nil && choice.ID != 0 {
		models.DeleteChoice(choice)
		c.JSON(http.StatusOK, gin.H{
			"message": "選擇題刪除成功",
		})
		return
	}
}

// DeleteOption 刪除選項
func DeleteOption(c *gin.Context) {
	data := struct {
		ID uint `json:"ID"`
	}{}
	c.BindJSON(&data)
	if opt, err := models.FindOption(data.ID); err == nil && opt.ID != 0 {
		models.DeleteOption(opt)
		c.JSON(http.StatusOK, gin.H{
			"message": "選項刪除成功",
		})
		return
	}
}
