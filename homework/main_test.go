package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"os/signal"
	"syscall"
	"testing"

	"NCNUOJBackend/homework/models"
	"NCNUOJBackend/homework/router"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/assert.v1"
)

var d struct {
	Token string `json:"token"`
}

func init() {
	gin.SetMode(gin.TestMode)
	models.Setup()
}

var sigs = make(chan os.Signal, 1)

// 創建題目
func TestProblemCreate(t *testing.T) {
	var data = []byte(`{
		"problem_name":       "接龍遊戲",
		"description":        "開始接龍",
		"input_description":  "567",
		"output_description": "789",
		"author":             123,
		"memory_limit":       123,
		"cpu_time":           123,
		"layer":              1,
		"sample_input":       ["123"],
		"sample_output":      ["456"],
		"tags_list":          ["簡單"]
	}`)
	r := router.SetupRouter()
	w := httptest.NewRecorder() // 取得 ResponseRecorder 物件
	req, _ := http.NewRequest("POST", "/api/v1/prob/createproblem", bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

// 列出所有題目id
func TestGetProblemList(t *testing.T) {
	r := router.SetupRouter()
	w := httptest.NewRecorder() // 取得 ResponseRecorder 物件
	req, _ := http.NewRequest("GET", "/api/v1/prob/list_of_problem", bytes.NewBuffer(make([]byte, 1000)))
	req.Header.Set("Authorization", "Bearer "+d.Token)
	r.ServeHTTP(w, req)
	body, _ := ioutil.ReadAll(w.Body)
	s := struct {
		ProblemsID []uint `json:"problemsid"`
	}{}
	json.Unmarshal(body, &s)
	assert.Equal(t, 1, len(s.ProblemsID))
}

// 用problemid 取得 problem的資訊
func TestProblemDetailByProbId(t *testing.T) {
	r := router.SetupRouter()
	w := httptest.NewRecorder() // 取得 ResponseRecorder 物件
	req, _ := http.NewRequest("GET", "/api/v1/prob/problem/1", bytes.NewBuffer(make([]byte, 1000)))
	req.Header.Set("Authorization", "Bearer "+d.Token)
	r.ServeHTTP(w, req)
	body, _ := ioutil.ReadAll(w.Body)
	s := struct {
		Problem_Name       string   `json:"problem_name"`
		Description        string   `json:"description"`
		Input_description  string   `json:"input_description"`
		Output_description string   `json:"output_description"`
		Author             uint     `json:"author"`
		Memory_limit       uint     `json:"memory_limit"`
		Cpu_time           uint     `json:"cpu_time"`
		Layer              uint8    `json:"layer"`
		Sample_input       []string `json:"sample_input"`
		Sample_output      []string `json:"sample_output"`
		Tags_list          []string `json:"tags_list"`
	}{}
	json.Unmarshal(body, &s)
	assert.Equal(t, "接龍遊戲", s.Problem_Name)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

}
func TestGetProblemByTagId(t *testing.T) {
	r := router.SetupRouter()
	w := httptest.NewRecorder() // 取得 ResponseRecorder 物件
	req, _ := http.NewRequest("GET", "/api/v1/prob/tag/1", bytes.NewBuffer(make([]byte, 1000)))
	req.Header.Set("Authorization", "Bearer "+d.Token)
	r.ServeHTTP(w, req)
	body, _ := ioutil.ReadAll(w.Body)
	s := struct {
		Problem_list   []string `json:"problem_names"`
		Problemid_list []uint   `json:"problem_ids"`
	}{}
	json.Unmarshal(body, &s)
	fmt.Println(s.Problemid_list)
	assert.Equal(t, 1, len(s.Problemid_list))
}

// 修改題目
func TestEditProblemByProblemId(t *testing.T) {
	var data = []byte(`{
		"problem_name":       "不是接龍遊戲",
		"description":        "開始接龍",
		"input_description":  "567",
		"output_description": "789",
		"author":             123,
		"memory_limit":       123,
		"cpu_time":           123,
		"layer":              1,
		"sample_input":       ["478"],
		"sample_output":      ["456"],
		"tags_list":          ["簡單"]
	}`)
	r := router.SetupRouter()
	w := httptest.NewRecorder() // 取得 ResponseRecorder 物件
	req, _ := http.NewRequest("PUT", "/api/v1/prob/edit/1", bytes.NewBuffer(data))
	req.Header.Set("Authorization", "Bearer "+d.Token)
	r.ServeHTTP(w, req)
	body, _ := ioutil.ReadAll(w.Body)
	s := struct {
		Message string `json:"message"`
	}{}
	json.Unmarshal(body, &s)
	assert.Equal(t, "題目更新成功", s.Message)
}

// 刪除題目id
func TestDeleteProblemById(t *testing.T) {
	r := router.SetupRouter()
	w := httptest.NewRecorder() // 取得 ResponseRecorder 物件
	req, _ := http.NewRequest("DELETE", "/api/v1/prob/delete/1", bytes.NewBuffer(make([]byte, 1000)))
	req.Header.Set("Authorization", "Bearer "+d.Token)
	r.ServeHTTP(w, req)
	body, _ := ioutil.ReadAll(w.Body)
	s := struct {
		Message string `json:"message"`
	}{}
	json.Unmarshal(body, &s)
	assert.Equal(t, "題目刪除成功", s.Message)
}
