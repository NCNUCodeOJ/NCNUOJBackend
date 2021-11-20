package view

import (
	"NCNUOJBackend/testPaper/models"
	"math/bits"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vincentinttsh/replace"
	"github.com/vincentinttsh/zero"
)

// AddQuestion 新增題目
func AddQuestion(c *gin.Context) {
	// 使用者傳過來的檔案格式(題目、出題者、範圍、出處)
	var questionData struct {
		Question   *string `json:"question"`
		AuthorID   *uint   `json:"authorID"`
		Layer      *uint   `json:"layer"`
		Source     *uint   `json:"source"`
		Difficulty *uint   `json:"difficulty"`
		Type       *uint   `json:"type"`
	}
	var question models.Question
	if err := c.BindJSON(&questionData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "選擇題未按照格式填寫",
		})
		return
	}
	// 如果有空值，則回傳 false
	if zero.IsZero(questionData) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "選擇題所有欄位不可為空值",
		})
		return
	}
	question.Question = *questionData.Question
	question.AuthorID = *questionData.AuthorID
	question.Layer = *questionData.Layer
	question.Source = *questionData.Source
	question.Difficulty = *questionData.Difficulty
	question.Type = *questionData.Type
	models.AddQuestion(&question)
	c.JSON(http.StatusOK, gin.H{
		"message": "新增題目成功",
	})
}

// AddAnswer 新增選項(答案)
func AddAnswer(c *gin.Context) {
	// 使用者傳過來的檔案格式(選項內容、是否為正確答案、對應的題目)
	var answerData struct {
		Content    *string `json:"content"`
		Correct    *bool   `json:"correct"`
		QuestionID *uint   `json:"questionID"`
		Sort       *uint   `json:"sort"`
	}
	var answer models.Answer
	if err := c.BindJSON(&answerData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "答案或選項未按照格式填寫",
		})
		return
	}
	// 如果有空值，則回傳 false
	if zero.IsZero(answerData) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "選項所有欄位不可為空值",
		})
		return
	}
	answer.Content = *answerData.Content
	answer.Correct = *answerData.Correct
	answer.QuestionID = *answerData.QuestionID
	answer.Sort = *answerData.Sort
	models.AddAnswer(&answer)
	c.JSON(http.StatusOK, gin.H{
		"message": "新增選項成功",
	})
}

// GetQuestion 查詢選擇題
func GetQuestion(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("questionID"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "路徑錯誤",
		})
		return
	}
	question, err := models.GetQuestion(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "查無此資料",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":         question.ID,
		"question":   question.Question,
		"authorID":   question.AuthorID,
		"layer":      question.Layer,
		"source":     question.Source,
		"difficulty": question.Difficulty,
		"type":       question.Type,
	})
}

// GetAllAnswers 透過 ID 取得測驗卷
func GetAllAnswers(c *gin.Context) {
	var allAnswerID []uint
	if answers, err := models.GetAllAnswers(); err == nil {
		for pos := range answers {
			allAnswerID = append(allAnswerID, answers[pos].ID)
		}
		c.JSON(http.StatusOK, gin.H{
			"answersID": allAnswerID,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "尚無內容",
		})
	}
}

// GetAnswer 查詢選項
func GetAnswer(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("answerID"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "路徑錯誤",
		})
		return
	}
	answer, err := models.GetAnswer(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "查無此資料",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":         answer.ID,
		"content":    answer.Content,
		"correct":    answer.Correct,
		"questionID": answer.QuestionID,
		"sort":       answer.Sort,
	})
}

// EditQuestion 修改選擇題
func EditQuestion(c *gin.Context) {
	id, err := strconv.ParseUint(c.Params.ByName("questionID"), 10, bits.UintSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "路徑錯誤",
		})
		return
	}
	var data struct {
		ID         *uint   `json:"ID"`
		Question   *string `json:"question"`
		AuthorID   *uint   `json:"authorID"`
		Layer      *uint   `json:"layer"`
		Source     *uint   `json:"source"`
		Difficulty *uint   `json:"difficulty"`
		Type       *uint   `json:"type"`
	}
	c.BindJSON(&data)
	question, err := models.GetQuestion(uint(id))
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "查無此資料",
		})
		return
	}
	replace.Replace(&question, &data)
	err = models.EditQuestion(&question)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "修改題目成功",
	})
}

// EditAnswer 修改答案/選項
func EditAnswer(c *gin.Context) {
	id, err := strconv.ParseUint(c.Params.ByName("answerID"), 10, bits.UintSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "路徑錯誤",
		})
		return
	}
	answerData := struct {
		ID         *uint   `json:"ID"`
		Content    *string `json:"content"`
		Correct    *bool   `json:"correct"`
		QuestionID *uint   `json:"questionID"`
		Sort       *uint   `json:"sort"`
	}{}
	c.BindJSON(&answerData)
	answer, err := models.GetAnswer(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "查無此資料",
		})
		return
	}
	replace.Replace(&answer, &answerData)
	err = models.EditAnswer(&answer)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "修改選項/答案成功",
	})
}

// DeleteQuestion 刪除選擇題
func DeleteQuestion(c *gin.Context) {
	id, err := strconv.ParseUint(c.Params.ByName("questionID"), 10, bits.UintSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "路徑錯誤",
		})
		return
	}
	question, err := models.GetQuestion(uint(id))
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "查無此資料",
		})
		return
	}
	err = models.DeleteQuestion(question)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "失敗",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "刪除問題成功",
	})
}

// DeleteAnswer 刪除答案
func DeleteAnswer(c *gin.Context) {
	id, err := strconv.ParseUint(c.Params.ByName("answerID"), 10, bits.UintSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "路徑錯誤",
		})
		return
	}
	answer, err := models.GetAnswer(uint(id))
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "查無此資料",
		})
		return
	}
	err = models.DeleteAnswer(answer)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "失敗",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "刪除答案/選項成功",
	})
}
