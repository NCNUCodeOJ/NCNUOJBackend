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

// AddTopic 新增
func AddTopic(c *gin.Context) {
	// 使用者傳過來的檔案格式(名稱、出卷者、對應的課堂、是否亂數出題)
	var data struct {
		Description *string `json:"description"`
		TestPaperID *uint   `json:"testPaperID"`
		Sort        *uint   `json:"sort"`
	}
	var topic models.Topic
	log.Println(data.TestPaperID == nil)
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "未按照格式填寫",
		})
		return
	}
	log.Println(data.TestPaperID == nil)
	// 如果有空值，則回傳 false
	if zero.IsZero(data) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "所有欄位不可為空值",
		})
		return
	}

	topic.Description = *data.Description
	topic.TestPaperID = *data.TestPaperID
	topic.Sort = *data.Sort
	models.AddTopic(&topic)
	c.JSON(http.StatusOK, gin.H{
		"message": "新增成功",
	})
}

// GetAllTopics 透過 ID 取得測驗卷
func GetAllTopics(c *gin.Context) {
	var allTopicID []uint
	if topics, err := models.GetAllTopics(); err == nil {
		for pos := range topics {
			allTopicID = append(allTopicID, topics[pos].ID)
		}
		c.JSON(http.StatusOK, gin.H{
			"topicsID": allTopicID,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "尚無內容",
		})
	}
}

// GetTopic 透過 ID 取得
func GetTopic(c *gin.Context) {
	// check redis
	id, err := strconv.Atoi(c.Params.ByName("topicID"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "系統錯誤",
		})
		return
	}
	topic, err := models.GetTopic(uint(id)) //這邊
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "查無資料",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":          topic.ID,
		"description": topic.Description,
		"testpaperID": topic.TestPaperID,
		"sort":        topic.Sort,
	})
}

// EditTopic 修改大題
func EditTopic(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("topicID"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "系統錯誤",
		})
		return
	}
	data := struct {
		Description *string `json:"description"`
		TestPaperID *uint   `json:"testPaperID"`
		Sort        *uint   `json:"sort"`
	}{}
	c.BindJSON(&data)
	topic, err := models.GetTopic(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "查無資料",
		})
		return
	}
	replace.Replace(&topic, &data)
	err = models.EditTopic(&topic)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "修改失敗",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "修改成功",
	})
}

// DeleteTopic 刪除
func DeleteTopic(c *gin.Context) {
	id, err := strconv.ParseUint(c.Params.ByName("topicID"), 10, bits.UintSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "系統錯誤",
		})
		return
	}
	topic, err := models.GetTopic(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "查無資料",
		})
		return
	}
	err = models.DeleteTopic(topic)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "刪除失敗",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "刪除成功",
	})
}
