package view

import (
	"NCNUOJBackend/testPaper/models"
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vincentinttsh/replace"
	"github.com/vincentinttsh/zero"
)

// SetTestPaper 新增測驗卷
func SetTestPaper(c *gin.Context) {
	// 使用者傳過來的檔案格式(測驗卷名稱、出卷者、對應的課堂、是否亂數出題)
	var data struct {
		Name     *string `json:"name"`
		AuthorID *uint   `json:"authorID"`
		ClassID  *uint   `json:"classID"`
		Random   *bool   `json:"random"`
	}
	var testpaper models.TestPaper
	log.Println(data.Name == nil)
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "測驗卷未按照格式填寫",
		})
		return
	}
	log.Println(data.Name == nil)
	// 如果有空值，則回傳 false
	if zero.IsZero(data) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "測驗卷所有欄位不可為空值",
		})
		return
	}

	testpaper.Name = *data.Name
	testpaper.AuthorID = *data.AuthorID
	testpaper.ClassID = *data.ClassID
	testpaper.Random = *data.Random
	models.AddTestPaper(&testpaper)
	c.JSON(http.StatusOK, gin.H{
		"message": "新增測驗卷成功",
	})
}

// Choices 新增測驗卷中的選擇題們
func ChoiceTP(c *gin.Context) {
	// 使用者傳過來的檔案格式(對應的課堂、對應的選擇題題目、排序)
	var data struct {
		ClassID  uint `json:"classID"`
		ChoiceID uint `json:"choiceID"`
		Sort     uint `json:"sort"`
	}
	var choiceTP models.ChoiceTP
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "選擇題們未按照格式填寫",
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
	choiceTP.ClassID = data.ClassID
	choiceTP.ChoiceID = data.ChoiceID
	choiceTP.Sort = data.Sort
	models.AddChoiceTP(&choiceTP)
	c.JSON(http.StatusOK, gin.H{
		"message": "新增選擇題們成功",
	})
}

// Clozes 新增測驗卷中的填充題們
func ClozeTP(c *gin.Context) {
	// 使用者傳過來的檔案格式(對應的課堂、對應的選擇題題目、排序)
	var data struct {
		ClassID uint `json:"classID"`
		ClozeID uint `json:"clozeID"`
		Sort    uint `json:"sort"`
	}
	var clozes models.ClozeTP
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "選擇題們未按照格式填寫",
		})
		return
	}
	clozes.ClassID = data.ClassID
	clozes.ClozeID = data.ClozeID
	clozes.Sort = data.Sort
	models.AddClozes(&clozes)
	c.JSON(http.StatusOK, gin.H{
		"message": "新增填充題們成功",
	})
}

// FindCloze 查詢選擇題
func FindTestPaper(c *gin.Context) {
	data := struct {
		ID uint `json:"ID"`
	}{}
	c.Bind(&data)
	if tp, err := models.FindTestPaper(data.ID); err == nil && tp.ID != 0 {
		c.JSON(http.StatusOK, gin.H{
			"id":       tp.ID,
			"name":     tp.Name,
			"authorID": tp.AuthorID,
			"classID":  tp.ClassID,
			"random":   tp.Random,
		})
		return
	}
}

// Edit 修改測驗卷
func EditTP(c *gin.Context) {
	data := struct {
		ID       *uint   `json:"ID"`
		Name     *string `gorm:"NOT NULL;"`
		AuthorID *uint
		ClassID  *uint
		Random   *bool
	}{}
	c.BindJSON(&data)
	if tp, err := models.FindTestPaper(*data.ID); err == nil && tp.ID != 0 {
		replace.Replace(&tp, &data)
		c.JSON(http.StatusOK, gin.H{
			"message": "修改成功",
		})
		return
	}
}

// DeleteTP 刪除測驗卷
func DeleteTP(c *gin.Context) {
	data := struct {
		ID uint `json:"ID"`
	}{}
	c.BindJSON(&data)
	if tp, err := models.FindTestPaper(data.ID); err == nil && tp.ID != 0 {
		models.DeleteTP(tp)
		c.JSON(http.StatusOK, gin.H{
			"message": "測驗卷刪除成功",
		})
		return
	}
}

// DeleteChoiceTP 刪除測驗卷中的選擇題
func DeleteChoiceTP(c *gin.Context) {
	data := struct {
		ID uint `json:"ID"`
	}{}
	c.BindJSON(&data)
	if choicetp, err := models.FindChoiceTP(data.ID); err == nil && choicetp.ID != 0 {
		models.DeleteChoiceTP(choicetp)
		c.JSON(http.StatusOK, gin.H{
			"message": "測驗卷中的選擇題刪除成功",
		})
		return
	}
}

// DeleteClozeTP 刪除測驗卷中的填充題
func DeleteClozeTP(c *gin.Context) {
	data := struct {
		ID uint `json:"ID"`
	}{}
	c.BindJSON(&data)
	if clozetp, err := models.FindClozeTP(data.ID); err == nil && clozetp.ID != 0 {
		models.DeleteClozeTP(clozetp)
		c.JSON(http.StatusOK, gin.H{
			"message": "測驗卷中的填充題刪除成功",
		})
		return
	}
}
