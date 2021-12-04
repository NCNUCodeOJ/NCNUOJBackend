package main

// go test main.go main_test.go

import (
	"NCNUOJBackend/testPaper/models"
	"NCNUOJBackend/testPaper/router"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/assert.v1"
)

var token = "Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6ZmFsc2UsImV4cCI6NDc5MTA4MjEyMywiaWQiOiI3MTI0MTMxNTQxOTcxMTA3ODYiLCJvcmlnX2lhdCI6MTYzNzQ4MjEyMywidXNlcm5hbWUiOiJ0ZXN0X3VzZXIifQ.pznOSok8X7qv6FSIihJnma_zEy70TerzOs0QDZOq_4RPYOKSEOOYTZ9-VLm2P9XRldS17-7QrLFwjjfXyCodtA"

func init() {
	gin.SetMode(gin.TestMode)
	models.Setup()
}

var sigs = make(chan os.Signal, 1)

func TestMain(t *testing.T) {
	start()
	// time.Sleep(10 * time.Second)
	end()
}

// // POST 測驗卷
// func TestCreateTestpaper(t *testing.T) {
// 	var testpaperData = []byte(`{
// 		"testpaper_name": "testpaper1",
// 		"author": 1,
// 		"class_id": 1,
// 		"random": false
// 	}`)
// 	r := router.SetupRouter()
// 	w := httptest.NewRecorder() // 取得 ResponseRecorder 物件
// 	req, _ := http.NewRequest("POST", "/api/private/v1/testpaper", bytes.NewBuffer(testpaperData))
// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set("Authorization", token)
// 	r.ServeHTTP(w, req)
// 	assert.Equal(t, http.StatusOK, w.Code)
// }

// // POST 大題
// func TestCreateTopic(t *testing.T) {
// 	var topicData = []byte(`{
// 		"description": "topic1",
// 		"testpaper_id": 1,
// 		"sort": 1
// 	}`)
// 	r := router.SetupRouter()
// 	w := httptest.NewRecorder() // 取得 ResponseRecorder 物件
// 	req, _ := http.NewRequest("POST", "/api/private/v1/testpaper/1/topic", bytes.NewBuffer(topicData))
// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set("Authorization", token)
// 	r.ServeHTTP(w, req)
// 	assert.Equal(t, http.StatusOK, w.Code)
// }

// POST 題目
func TestCreateQuestion(t *testing.T) {
	var questionData = []byte(`{
		"question": "question1",
		"author": 1,
		"layer": 1,
		"source": "source",
		"difficulty": 1,
		"type": 1,
		"option": [
			{
				"content": "content1",
				"answer": true,
				"question_id": 1,
				"sort": 1
			},
			{
				"content": "content2",
				"answer": true,
				"question_id": 1,
				"sort": 2
			}
		  ]
	}`)
	r := router.SetupRouter()
	w := httptest.NewRecorder() // 取得 ResponseRecorder 物件
	req, _ := http.NewRequest("POST", "/api/v1/question", bytes.NewBuffer(questionData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

// // GET 所有測驗卷
// func TestListTestPapers(t *testing.T) {
// 	r := router.SetupRouter()
// 	w := httptest.NewRecorder() // 取得 ResponseRecorder 物件
// 	req, _ := http.NewRequest("GET", "/api/private/v1/testpaper", bytes.NewBuffer(make([]byte, 1000)))
// 	req.Header.Set("Authorization", token)
// 	r.ServeHTTP(w, req)
// 	assert.Equal(t, http.StatusOK, w.Code)
// 	body, _ := ioutil.ReadAll(w.Body)
// 	s := struct {
// 		TestpapersID []uint `json:"testpapers_id"`
// 	}{}
// 	json.Unmarshal(body, &s)
// 	r.ServeHTTP(w, req)
// }

// // GET 測驗卷
// func TestGetTestpaperByID(t *testing.T) {
// 	r := router.SetupRouter()
// 	// 取得 ResponseRecorder 物件，用來記錄 response 狀態
// 	w := httptest.NewRecorder()
// 	req, _ := http.NewRequest("GET", "/api/private/v1/testpaper/1", bytes.NewBuffer(make([]byte, 1000)))
// 	req.Header.Set("Authorization", token)
// 	// gin.Engine.ServerHttp 實作 http.Handler 介面，用來處理 HTTP 請求及回應
// 	r.ServeHTTP(w, req)
// 	body, _ := ioutil.ReadAll(w.Body)
// 	s := struct {
// 		TestPaperName string `json:"testpaper_name"`
// 		Author        uint   `json:"author"`
// 		ClassID       uint   `json:"class_id"`
// 		Random        bool   `json:"random"`
// 	}{}
// 	json.Unmarshal(body, &s)
// 	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
// 	assert.Equal(t, http.StatusOK, w.Code)
// }

// // GET 所有大題
// func TestListTopics(t *testing.T) {
// 	r := router.SetupRouter()
// 	w := httptest.NewRecorder() // 取得 ResponseRecorder 物件
// 	req, _ := http.NewRequest("GET", "/api/private/v1/testpaper/1/topic", bytes.NewBuffer(make([]byte, 1000)))
// 	req.Header.Set("Authorization", token)
// 	r.ServeHTTP(w, req)
// 	assert.Equal(t, http.StatusOK, w.Code)
// 	body, _ := ioutil.ReadAll(w.Body)
// 	s := struct {
// 		TopicID []uint `json:"topics_id"`
// 	}{}
// 	json.Unmarshal(body, &s)
// 	r.ServeHTTP(w, req)
// }

// // Get 大題
// func TestGetTopic(t *testing.T) {
// 	r := router.SetupRouter()
// 	w := httptest.NewRecorder() // 取得 ResponseRecorder 物件
// 	req, _ := http.NewRequest("GET", "/api/private/v1/testpaper/1/topic/1", bytes.NewBuffer(make([]byte, 1000)))
// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set("Authorization", token)
// 	r.ServeHTTP(w, req)
// 	assert.Equal(t, http.StatusOK, w.Code)
// }

// Get 所有題目
func TestListQuestions(t *testing.T) {
	r := router.SetupRouter()
	w := httptest.NewRecorder() // 取得 ResponseRecorder 物件
	req, _ := http.NewRequest("GET", "/api/v1/question", bytes.NewBuffer(make([]byte, 1000)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
} // }

// Get 題目
func TestGetQuestion(t *testing.T) {
	r := router.SetupRouter()
	w := httptest.NewRecorder() // 取得 ResponseRecorder 物件
	req, _ := http.NewRequest("GET", "/api/v1/question/1", bytes.NewBuffer(make([]byte, 1000)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

// // PATCH 測驗卷
// func TestUpdateTestpaper(t *testing.T) {
// 	var testpaperPatchData = []byte(`{
// 		"testpaper_name": "testpaper1patchtest",
// 		"author": 1patch,
// 		"class_id": 1patch,
// 		"random": false
// 	}`)
// 	r := router.SetupRouter()
// 	w := httptest.NewRecorder() // 取得 ResponseRecorder 物件
// 	req, _ := http.NewRequest("PATCH", "/api/private/v1/testpaper/1", bytes.NewBuffer(testpaperPatchData))
// 	req.Header.Set("Authorization", token)
// 	r.ServeHTTP(w, req)
// 	body, _ := ioutil.ReadAll(w.Body)
// 	s := struct {
// 		Message string `json:"message"`
// 	}{}
// 	json.Unmarshal(body, &s)
// 	assert.Equal(t, http.StatusOK, w.Code)
// }

// // PATCH 大題
// func TestUpdateTopic(t *testing.T) {
// 	var topicData = []byte(`{
// 		"description": "topicpatchtest",
// 		"testpaper_id": 1,
// 		"sort": 1
// 	}`)
// 	r := router.SetupRouter()
// 	w := httptest.NewRecorder() // 取得 ResponseRecorder 物件
// 	req, _ := http.NewRequest("PATCH", "/api/private/v1/testpaper/1/topic/1", bytes.NewBuffer(topicData))
// 	req.Header.Set("Authorization", token)
// 	r.ServeHTTP(w, req)
// 	body, _ := ioutil.ReadAll(w.Body)
// 	s := struct {
// 		Message string `json:"message"`
// 	}{}
// 	json.Unmarshal(body, &s)
// }

// PATCH 題目
func TestUpdateQuestion(t *testing.T) {
	var questionData = []byte(`{
		"question": "question1patchtest",
		"author": 1,
		"layer": 1,
		"source": 1,
		"difficulty": 1,
		"type": 1
	}`)
	r := router.SetupRouter()
	w := httptest.NewRecorder() // 取得 ResponseRecorder 物件
	req, _ := http.NewRequest("PATCH", "/api/v1/question/1", bytes.NewBuffer(questionData))
	req.Header.Set("Authorization", token)
	r.ServeHTTP(w, req)
	body, _ := ioutil.ReadAll(w.Body)
	s := struct {
		Message string `json:"message"`
	}{}
	json.Unmarshal(body, &s)
}

// // Delete 測驗卷
// func TestDeleteTestpaper(t *testing.T) {
// 	r := router.SetupRouter()
// 	w := httptest.NewRecorder() // 取得 ResponseRecorder 物件
// 	req, _ := http.NewRequest("DELETE", "/api/private/v1/testpaper/1", bytes.NewBuffer(make([]byte, 1000)))
// 	req.Header.Set("Authorization", token)
// 	r.ServeHTTP(w, req)
// 	body, _ := ioutil.ReadAll(w.Body)
// 	s := struct {
// 		Message string `json:"message"`
// 	}{}
// 	json.Unmarshal(body, &s)
// 	assert.Equal(t, http.StatusOK, w.Code)
// }

// // // Delete 大題
// func TestDeleteTopic(t *testing.T) {
// 	r := router.SetupRouter()
// 	w := httptest.NewRecorder() // 取得 ResponseRecorder 物件
// 	req, _ := http.NewRequest("DELETE", "/api/private/v1/testpaper/1/topic/1", bytes.NewBuffer(make([]byte, 1000)))
// 	req.Header.Set("Authorization", token)
// 	r.ServeHTTP(w, req)
// 	body, _ := ioutil.ReadAll(w.Body)
// 	s := struct {
// 		Message string `json:"message"`
// 	}{}
// 	json.Unmarshal(body, &s)
// 	assert.Equal(t, http.StatusOK, w.Code)
// }

// Delete 題目
func TestDeleteQuestion(t *testing.T) {
	r := router.SetupRouter()
	w := httptest.NewRecorder() // 取得 ResponseRecorder 物件
	req, _ := http.NewRequest("DELETE", "/api/v1/question/1", bytes.NewBuffer(make([]byte, 1000)))
	req.Header.Set("Authorization", token)
	r.ServeHTTP(w, req)
	body, _ := ioutil.ReadAll(w.Body)
	s := struct {
		Message string `json:"message"`
	}{}
	json.Unmarshal(body, &s)
	assert.Equal(t, http.StatusOK, w.Code)
}
