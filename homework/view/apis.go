package view

import (
	"NCNUOJBackend/homework/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vincentinttsh/zero"
)

//CreateProblem 創建題目
func CreateProblem(c *gin.Context) {
	data := struct {
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

	var problem models.Problem

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "未按照格式填寫",
		})
		return
	}
	if zero.IsZero(data) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "未按照格式填寫",
		})
		return
	}
	if len(data.Sample_input) != len(data.Sample_output) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "未按照格式填寫",
		})
		return
	}
	problem.ProblemName = data.Problem_Name
	problem.Description = data.Description
	problem.InputDescription = data.Input_description
	problem.OutputDescripton = data.Output_description
	problem.Author = data.Author
	problem.MemoryLimit = data.Memory_limit
	problem.Cputime = data.Cpu_time
	problem.Layer = data.Layer
	// replace.Replace(&problem, &data)
	if err := models.AddProblem(&problem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "題目創建失敗",
		})
	}

	for pos := range data.Tags_list {
		var tag models.Tag
		var tag2table models.Tag2Table

		if t, err := models.CheckTag(data.Tags_list[pos]); err == nil && t.ID != 0 {
			tag2table.ProblemId = problem.ID
			tag2table.TagId = t.ID
			models.AddTag2Table(&tag2table)
		} else {
			tag.Name = data.Tags_list[pos]
			models.AddTag(&tag)

			tag2table.ProblemId = problem.ID
			tag2table.TagId = tag.ID
			models.AddTag2Table(&tag2table)
		}
	}

	for pos := range data.Sample_input {
		var sample models.Sample

		sample.Input = data.Sample_input[pos]
		sample.Output = data.Sample_output[pos]
		sample.ProblemId = problem.ID

		models.AddSample(&sample)

	}

	c.JSON(http.StatusOK, gin.H{
		"message": "題目創建成功",
	})
}

//GetProblemByProblemId 用 problem id 找 problem
func GetProblemByProblemId(c *gin.Context) {
	Id, err := strconv.Atoi(c.Params.ByName("probid"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "系統錯誤",
		})
	}
	var ProblemId uint = uint(Id)
	if zero.IsZero(Id) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "未按照格式填寫",
		})
		return
	}
	if u, err := models.ProblemDetailByProblemId(ProblemId); err == nil && u.ID != 0 {
		var SampleIn []string
		var SampleOut []string
		var tags []string

		if s, err := models.SampleDetailByProblemId(ProblemId); err == nil && len(s) != 0 {
			for pos := range s {
				SampleIn = append(SampleIn, s[pos].Input)
				SampleOut = append(SampleOut, s[pos].Output)
			}
		}

		if t, err := models.TagDetailByProblemId(ProblemId); err == nil && len(t) != 0 {
			for pos := range t {
				if t1, err := models.TagDetailByTagId(t[pos].TagId); err == nil && t1.ID != 0 {
					tags = append(tags, t1.Name)
				}
			}

		}
		c.JSON(http.StatusOK, gin.H{
			"problem_id":         u.ID,
			"problem_name":       u.ProblemName,
			"description":        u.Description,
			"input_description":  u.InputDescription,
			"output_description": u.OutputDescripton,
			"author":             u.Author,
			"memory_limit":       u.MemoryLimit,
			"cpu_time":           u.Cputime,
			"layer":              u.Layer,
			"sample_input":       SampleIn,
			"sample_output":      SampleOut,
			"tags_list":          tags,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "無此題目",
		})
	}
}

//GetProblemByTagId 用 tag id 找有哪些 problem
func GetProblemByTagId(c *gin.Context) {
	var problem_list []string
	var problemid_list []uint
	Id, err := strconv.Atoi(c.Params.ByName("tagid"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "系統錯誤",
		})
	}
	var TagId uint = uint(Id)
	if zero.IsZero(TagId) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "未按照格式填寫",
		})
		return
	}
	if t, err := models.ProblemDetailByTagId(TagId); err == nil && len(t) != 0 {

		for pos := range t {

			if p, err := models.ProblemDetailByProblemId(t[pos].ProblemId); err == nil && p.ID != 0 {
				problem_list = append(problem_list, p.ProblemName)
				problemid_list = append(problemid_list, p.ID)

			}
		}
		c.JSON(http.StatusOK, gin.H{
			"problem_names": problem_list,
			"problem_ids":   problemid_list,
		})

		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "還未有題目使用這個標籤",
		})
	}

}

//EditProblemByProblemId 編輯題目
func EditProblemByProblemId(c *gin.Context) {
	Id, err := strconv.Atoi(c.Params.ByName("probid"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "系統錯誤",
		})
	}
	var Problem_id uint = uint(Id)
	data := struct {
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

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "未按照格式填寫",
		})
		return
	}
	if zero.IsZero(data) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "未按照格式填寫",
		})
		return
	}

	// sample input ouput 的 item 數量一定要一樣
	if len(data.Sample_input) != len(data.Sample_output) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "未按照格式填寫",
		})
		return
	}
	// 更新problem
	var problem models.Problem
	problem.ID = Problem_id
	problem.ProblemName = data.Problem_Name
	problem.Description = data.Description
	problem.InputDescription = data.Input_description
	problem.OutputDescripton = data.Output_description
	problem.Author = data.Author
	problem.MemoryLimit = data.Memory_limit
	problem.Cputime = data.Cpu_time
	problem.Layer = data.Layer

	models.UpdateProblem(&problem)

	// 更新 tag
	var old_tags []models.Tag   //資料庫的舊資料
	var old_tagid []uint        //要被刪掉的舊資料
	var repeat_tagname []string // 重複的 tag name
	if t, err := models.TagDetailByProblemId(Problem_id); err == nil && len(t) != 0 {
		//先找舊的tag資料
		for pos := range t {
			if tag_name, err := models.TagDetailByTagId(t[pos].TagId); err == nil && tag_name.ID != 0 {
				old_tags = append(old_tags, tag_name)
			}
		}
		//先移除list裡重複的item，再用剩下的item刪掉資料庫的資料
		for o_pos := range old_tags {
			turn := 0
			for pos := range data.Tags_list {
				if old_tags[o_pos].Name == data.Tags_list[pos] {
					turn = 1
					repeat_tagname = append(repeat_tagname, old_tags[o_pos].Name)
					break
				}
			}
			if turn == 0 {
				old_tagid = append(old_tagid, old_tags[o_pos].ID)
			}
		}
		//刪掉更新後沒有的資料
		for pos := range old_tagid {
			models.EditForDeleteTag2Table(old_tagid[pos], Problem_id)
		}
	}

	for pos := range data.Tags_list {
		var tag models.Tag
		var tag2table models.Tag2Table
		var turn uint // 開關 做跳過的動作
		turn = 1
		//跳過已經有的資料
		for r_pos := range repeat_tagname {
			if data.Tags_list[pos] == repeat_tagname[r_pos] {
				turn = 0
			}
		}
		if turn == 0 {
			continue
		}
		//確認有沒有這個tag 有的話直接用tag id 去串接 problem id ，沒有就創建一個tag ，再用這個tag id 去串
		if t, err := models.CheckTag(data.Tags_list[pos]); err == nil && t.ID != 0 {
			tag2table.ProblemId = Problem_id
			tag2table.TagId = t.ID
			models.AddTag2Table(&tag2table)
		} else {
			tag.Name = data.Tags_list[pos]
			models.AddTag(&tag)

			tag2table.ProblemId = Problem_id
			tag2table.TagId = tag.ID
			models.AddTag2Table(&tag2table)
		}
	}

	var old_sample []models.Sample    //資料庫裡的舊 sample
	var repeat_sample []models.Sample //重複的 sample
	var old_sampleid []uint           //要被刪掉的舊 sample
	if s, err := models.SampleDetailByProblemId(Problem_id); err == nil && len(s) != 0 {
		for pos := range s {
			old_sample = append(old_sample, s[pos])
		}

		//先移除list裡重複的item，再用剩下的item刪掉資料庫的資料
		for o_pos := range old_sample {
			turn := 0
			for pos := range data.Sample_input {
				if old_sample[o_pos].Input == data.Sample_input[pos] && old_sample[o_pos].Output == data.Sample_output[pos] {
					turn = 1
					repeat_sample = append(repeat_sample, old_sample[o_pos])
					break
				}
			}
			// 放到 要被刪掉的資料列表裡
			if turn == 0 {
				old_sampleid = append(old_sampleid, old_sample[o_pos].ID)
			}
		}
		//刪掉更新後已經沒有用的資料
		for pos := range old_sampleid {
			models.EditForDeleteSample(old_sampleid[pos], Problem_id)
		}
	}

	for pos := range data.Sample_input {
		var turn uint //開關 略過的動作
		turn = 0
		// 略過重複的 sample
		for o_pos := range repeat_sample {
			if repeat_sample[o_pos].Input == data.Sample_input[pos] && repeat_sample[o_pos].Output == data.Sample_output[pos] {

				turn = 1
				break
			}
		}

		if turn == 1 {
			continue
		}
		// 創建新的 sample
		var sample models.Sample

		sample.Input = data.Sample_input[pos]
		sample.Output = data.Sample_output[pos]
		sample.ProblemId = Problem_id
		models.AddSample(&sample)

	}

	c.JSON(http.StatusOK, gin.H{
		"message": "題目更新成功",
	})
}

//DeleteProblemById
func DeleteProblemById(c *gin.Context) {
	Id, err := strconv.Atoi(c.Params.ByName("probid"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "系統錯誤",
		})
	}
	var ProblemId uint = uint(Id)
	if zero.IsZero(ProblemId) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "未按照格式填寫",
		})
		return
	}
	if err := models.DeleteProblem(ProblemId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "題目刪除失敗",
		})
		return
	}
	if err := models.DeleteSample(ProblemId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "題目刪除失敗",
		})
		return
	}
	if err := models.DeleteTag2TableByProb(ProblemId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "題目刪除失敗",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "題目刪除成功",
	})
}

//GetProblemList 列出所有problem id
func GetProblemList(c *gin.Context) {
	var problemid_list []uint
	if problems, err := models.ListProblem(); err == nil {
		for pos := range problems {
			problemid_list = append(problemid_list, problems[pos].ID)
		}
		c.JSON(http.StatusOK, gin.H{
			"problemsid": problemid_list,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "尚無題目",
		})
	}

}
