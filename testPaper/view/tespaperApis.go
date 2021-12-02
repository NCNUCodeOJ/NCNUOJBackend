package view

import (
	"NCNUOJBackend/testPaper/models"
	"log"
	"math/bits"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vincentinttsh/replace"
	"github.com/vincentinttsh/zero"
)

// AddTestPaper 新增測驗卷
func AddTestPaper(c *gin.Context) {
	var testpaper models.TestPaper
	userID := c.MustGet("userID").(uint)
	// 使用者傳過來的檔案格式(測驗卷名稱、出卷者、對應的課堂、是否亂數出題)
	var data struct {
		TestPaperName *string `json:"testpaperName"`
		AuthorID      *uint   `json:"authorID"`
		ClassID       *uint   `json:"classID"`
		Random        *bool   `json:"random"`
	}
	log.Println(data.TestPaperName == nil)
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "未按照格式填寫或未使用 json",
		})
		return
	}
	log.Println(data.TestPaperName == nil)
	// 如果有空值，則回傳 false
	if zero.IsZero(data) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "未填寫完成",
		})
		return
	}
	testpaper.TestPaperName = *data.TestPaperName
	testpaper.AuthorID = userID
	testpaper.ClassID = *data.ClassID
	testpaper.Random = *data.Random
	models.AddTestPaper(&testpaper)
	c.JSON(http.StatusOK, gin.H{
		"message": "新增成功",
	})
}

// GetAllTestPapers 取得全部測驗卷的 ID
func GetAllTestPapers(c *gin.Context) {
	var allTestpaperID []uint
	if testpapers, err := models.GetAllTestPapers(); err == nil {
		for pos := range testpapers {
			allTestpaperID = append(allTestpaperID, testpapers[pos].ID)
		}
		c.JSON(http.StatusOK, gin.H{
			"testpapersID": allTestpaperID,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "查無資料",
		})
	}
}

// GetTestPaper 透過 ID 取得測驗卷
func GetTestPaper(c *gin.Context) {
	// check redis
	// ParseUint convert strings to values
	id, err := strconv.Atoi(c.Params.ByName("testpaperID"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "系統錯誤",
		})
		return
	}
	testpaper, err := models.GetTestPaper(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "查無資料",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":            testpaper.ID,
		"testpaperName": testpaper.TestPaperName,
		"authorID":      testpaper.AuthorID,
		"classID":       testpaper.ClassID,
		"random":        testpaper.Random,
	})
}

// EditTestPaper 修改測驗卷
func EditTestPaper(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("testpaperID"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "系統錯誤",
		})
		return
	}
	var data struct {
		TestPaperName *string `json:"testpaperName"`
		AuthorID      *uint   `json:"authorID"`
		ClassID       *uint   `json:"classID"`
		Random        *bool   `json:"random"`
	}
	c.BindJSON(&data)
	testpaper, err := models.GetTestPaper(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "查無資料",
		})
		return
	}
	replace.Replace(&testpaper, &data)
	err = models.EditTestPaper(&testpaper)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "未填寫完成",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "修改成功",
	})
}

// DeleteTestPaper 刪除測驗卷
func DeleteTestPaper(c *gin.Context) {
	id, err := strconv.ParseUint(c.Params.ByName("testpaperID"), 10, bits.UintSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "系統錯誤",
		})
		return
	}
	testpaper, err := models.GetTestPaper(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "查無資料",
		})
		return
	}
	err = models.DeleteTestPaper(testpaper)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "刪除失敗",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "刪除成功",
	})
}
