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
	"os/signal"
	"syscall"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/assert.v1"
)

func init() {
	gin.SetMode(gin.TestMode)
	models.Setup()
}

var d struct {
	Token string `json:"token"`
}

var sigs = make(chan os.Signal, 1)

func TestMain(t *testing.T) {
	start()
	time.Sleep(10 * time.Second)
	end()
}

// POST 測驗卷
func TestAddTestpaper(t *testing.T) {
	var testpaperData = []byte(`{
		"testpapername": "testpaper1",
		"authorID": 1,
		"classID": 1,
		"random": false
	}`)
	r := router.SetupRouter()
	w := httptest.NewRecorder() // 取得 ResponseRecorder 物件
	req, _ := http.NewRequest("POST", "/api/v1/testpaper", bytes.NewBuffer(testpaperData))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

// POST 大題
func TestAddTopic(t *testing.T) {
	var topicData = []byte(`{
		"description": "topic1",
		"testpaperID": 1,
		"sort": 1
	}`)
	r := router.SetupRouter()
	w := httptest.NewRecorder() // 取得 ResponseRecorder 物件
	req, _ := http.NewRequest("POST", "/api/v1/testpaper/1/topic", bytes.NewBuffer(topicData))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

// POST 題目
func TestAddQuestion(t *testing.T) {
	var questionData = []byte(`{
		"question": "question1",
		"authorID": 1,
		"layer": 1,
		"source": 1,
		"difficulty": 1,
		"type": 1
	}`)
	r := router.SetupRouter()
	w := httptest.NewRecorder() // 取得 ResponseRecorder 物件
	req, _ := http.NewRequest("POST", "/api/v1/question", bytes.NewBuffer(questionData))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

// POST 選項/答案
func TestAddAnswer(t *testing.T) {
	var answerData = []byte(`{
		"content": "answer1",
		"correct": false,
		"questionID": 1,
		"sort": 1
	}`)
	r := router.SetupRouter()
	w := httptest.NewRecorder() // 取得 ResponseRecorder 物件
	req, _ := http.NewRequest("POST", "/api/v1/question/1/answer", bytes.NewBuffer(answerData))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

// GET 所有測驗卷
func TestGetAllTestpapers(t *testing.T) {
	r := router.SetupRouter()
	w := httptest.NewRecorder() // 取得 ResponseRecorder 物件
	req, _ := http.NewRequest("GET", "/api/v1/testpaper", bytes.NewBuffer(make([]byte, 1000)))
	req.Header.Set("Authorization", "Bearer "+d.Token)
	r.ServeHTTP(w, req)
	body, _ := ioutil.ReadAll(w.Body)
	s := struct {
		TestpapersID []uint `json:"testpapersID"`
	}{}
	json.Unmarshal(body, &s)
	// assert.Equal(t, 1, len(s.TestpapersID))
}

// GET 測驗卷
func TestGetTestpaper(t *testing.T) {
	r := router.SetupRouter()
	// 取得 ResponseRecorder 物件，用來記錄 response 狀態
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/testpaper/1", bytes.NewBuffer(make([]byte, 1000)))
	req.Header.Set("Authorization", "Bearer "+d.Token)
	// gin.Engine.ServerHttp 實作 http.Handler 介面，用來處理 HTTP 請求及回應
	r.ServeHTTP(w, req)
	body, _ := ioutil.ReadAll(w.Body)
	s := struct {
		Problem_Name       string `json:"testpapername"`
		Description        string `json:"authorID"`
		Input_description  string `json:"classID"`
		Output_description string `json:"random"`
	}{}
	json.Unmarshal(body, &s)
	// assert.Equal(t, "接龍遊戲", s.Problem_Name)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
}

// GET 所有大題
func TestGetAllTopics(t *testing.T) {
	r := router.SetupRouter()
	w := httptest.NewRecorder() // 取得 ResponseRecorder 物件
	req, _ := http.NewRequest("GET", "/api/v1/testpaper/1/topic", bytes.NewBuffer(make([]byte, 1000)))
	req.Header.Set("Authorization", "Bearer "+d.Token)
	r.ServeHTTP(w, req)
	body, _ := ioutil.ReadAll(w.Body)
	s := struct {
		TopicID []uint `json:"topicID"`
	}{}
	json.Unmarshal(body, &s)
	// assert.Equal(t, 1, len(s.TopicID))
}

// Get 大題
func TestGetTopic(t *testing.T) {
	r := router.SetupRouter()
	w := httptest.NewRecorder() // 取得 ResponseRecorder 物件
	req, _ := http.NewRequest("GET", "/api/v1/testpaper/1/topic/1", bytes.NewBuffer(make([]byte, 1000)))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	body, _ := ioutil.ReadAll(w.Body)
	json.Unmarshal(body, &d)
}

// Get 題目
func TestGetQuestion(t *testing.T) {
	r := router.SetupRouter()
	w := httptest.NewRecorder() // 取得 ResponseRecorder 物件
	req, _ := http.NewRequest("GET", "/api/v1/question/1", bytes.NewBuffer(make([]byte, 1000)))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	body, _ := ioutil.ReadAll(w.Body)
	json.Unmarshal(body, &d)
}

// Get 選項/答案
func TestGetAnswer(t *testing.T) {
	r := router.SetupRouter()
	w := httptest.NewRecorder() // 取得 ResponseRecorder 物件
	req, _ := http.NewRequest("GET", "/api/v1/question/1/answer/1", bytes.NewBuffer(make([]byte, 1000)))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	body, _ := ioutil.ReadAll(w.Body)
	json.Unmarshal(body, &d)
}

// PATCH 測驗卷
func TestEditTestpaper(t *testing.T) {
	var testpaperPatchData = []byte(`{
		"testpapername": "testpaper1patchtest",
		"authorID": 1patch,
		"classID": 1patch,
		"random": false
	}`)
	r := router.SetupRouter()
	w := httptest.NewRecorder() // 取得 ResponseRecorder 物件
	req, _ := http.NewRequest("PATCH", "/api/v1/testpaper/1", bytes.NewBuffer(testpaperPatchData))
	req.Header.Set("Authorization", "Bearer "+d.Token)
	r.ServeHTTP(w, req)
	body, _ := ioutil.ReadAll(w.Body)
	s := struct {
		Message string `json:"message"`
	}{}
	json.Unmarshal(body, &s)
}

// PATCH 大題
func TestEditTopic(t *testing.T) {
	var topicData = []byte(`{
		"description": "topicpatchtest",
		"testpaperID": 1,
		"sort": 1
	}`)
	r := router.SetupRouter()
	w := httptest.NewRecorder() // 取得 ResponseRecorder 物件
	req, _ := http.NewRequest("PATCH", "/api/v1/testpaper/1/topic/1", bytes.NewBuffer(topicData))
	req.Header.Set("Authorization", "Bearer "+d.Token)
	r.ServeHTTP(w, req)
	body, _ := ioutil.ReadAll(w.Body)
	s := struct {
		Message string `json:"message"`
	}{}
	json.Unmarshal(body, &s)
}

// PATCH 題目
func TestEditQuestion(t *testing.T) {
	var questionData = []byte(`{
		"question": "question1patchtest",
		"authorID": 1,
		"layer": 1,
		"source": 1,
		"difficulty": 1,
		"type": 1
	}`)
	r := router.SetupRouter()
	w := httptest.NewRecorder() // 取得 ResponseRecorder 物件
	req, _ := http.NewRequest("PATCH", "/api/v1/question/1", bytes.NewBuffer(questionData))
	req.Header.Set("Authorization", "Bearer "+d.Token)
	r.ServeHTTP(w, req)
	body, _ := ioutil.ReadAll(w.Body)
	s := struct {
		Message string `json:"message"`
	}{}
	json.Unmarshal(body, &s)
}

// PATCH 選項/答案
func TestEditAnswer(t *testing.T) {
	var answerData = []byte(`{
		"content": "answer1patchtest",
		"correct": false,
		"questionID": 1,
		"sort": 1
	}`)
	r := router.SetupRouter()
	w := httptest.NewRecorder() // 取得 ResponseRecorder 物件
	req, _ := http.NewRequest("PATCH", "/api/v1/question/1/answer/1", bytes.NewBuffer(answerData))
	req.Header.Set("Authorization", "Bearer "+d.Token)
	r.ServeHTTP(w, req)
	body, _ := ioutil.ReadAll(w.Body)
	s := struct {
		Message string `json:"message"`
	}{}
	json.Unmarshal(body, &s)
}

// Delete 測驗卷
func TestDeleteTestpaper(t *testing.T) {
	r := router.SetupRouter()
	w := httptest.NewRecorder() // 取得 ResponseRecorder 物件
	req, _ := http.NewRequest("DELETE", "/api/v1/testpaper/1", bytes.NewBuffer(make([]byte, 1000)))
	req.Header.Set("Authorization", "Bearer "+d.Token)
	r.ServeHTTP(w, req)
	body, _ := ioutil.ReadAll(w.Body)
	s := struct {
		Message string `json:"message"`
	}{}
	json.Unmarshal(body, &s)
	assert.Equal(t, http.StatusOK, w.Code)
}

// Delete 大題
func TestDeleteTopic(t *testing.T) {
	r := router.SetupRouter()
	w := httptest.NewRecorder() // 取得 ResponseRecorder 物件
	req, _ := http.NewRequest("DELETE", "/api/v1/testpaper/1/topic/1", bytes.NewBuffer(make([]byte, 1000)))
	req.Header.Set("Authorization", "Bearer "+d.Token)
	r.ServeHTTP(w, req)
	body, _ := ioutil.ReadAll(w.Body)
	s := struct {
		Message string `json:"message"`
	}{}
	json.Unmarshal(body, &s)
	assert.Equal(t, http.StatusOK, w.Code)
}

// Delete 題目
func TestDeleteQuestion(t *testing.T) {
	r := router.SetupRouter()
	w := httptest.NewRecorder() // 取得 ResponseRecorder 物件
	req, _ := http.NewRequest("DELETE", "/api/v1/question/1", bytes.NewBuffer(make([]byte, 1000)))
	req.Header.Set("Authorization", "Bearer "+d.Token)
	r.ServeHTTP(w, req)
	body, _ := ioutil.ReadAll(w.Body)
	s := struct {
		Message string `json:"message"`
	}{}
	json.Unmarshal(body, &s)
	assert.Equal(t, http.StatusOK, w.Code)
}

// Delete 選項/答案
func TestDeleteAnswer(t *testing.T) {
	r := router.SetupRouter()
	w := httptest.NewRecorder() // 取得 ResponseRecorder 物件
	req, _ := http.NewRequest("DELETE", "/api/v1/question/1/answer/1", bytes.NewBuffer(make([]byte, 1000)))
	req.Header.Set("Authorization", "Bearer "+d.Token)
	r.ServeHTTP(w, req)
	body, _ := ioutil.ReadAll(w.Body)
	s := struct {
		Message string `json:"message"`
	}{}
	json.Unmarshal(body, &s)
	assert.Equal(t, http.StatusOK, w.Code)
}

// TestAddQuestionTopic 測試 新增中間表
// func TestAddQuestionTopic(t *testing.T) {
// 	var questionTopicData = []byte(`{
// 		"topicID": "topic1",
// 		"distribution": 0.5,
// 		"sort": 1,
// 		"random": 1,
// 		"type": 1
// 	}`)
// 	r := router.SetupRouter()
// 	w := httptest.NewRecorder() // 取得 ResponseRecorder 物件
// 	req, _ := http.NewRequest("POST", "/api/v1/questionTopic", bytes.NewBuffer(questionTopicData))
// 	req.Header.Set("Content-Type", "application/json")
// 	r.ServeHTTP(w, req)
// 	assert.Equal(t, http.StatusOK, w.Code)
// }
