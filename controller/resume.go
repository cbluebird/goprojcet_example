package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"resume/db"
	"resume/db/model"
	"resume/utility"
	"strconv"
)

type SummaryOFResumes struct {
	Title    string
	Date     string
	Resumeid int
}

// 添加建立
func Add(c *gin.Context) {
	var data model.Resume
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utility.ResponseBadRequest(c)
		return
	}
	err = model.GetResumeByResume_nameAndUserid(data.Title, data.Userid)
	if err == nil {
		log.Println(err)
		utility.Response(404, "resume name repeat", nil, c)
		return
	}
	data.Date = utility.GetData()
	var last model.Resume
	db.DB.Last(&last)
	data.Resumeid = last.Resumeid + 1
	err = model.AddResume(&data)
	if err != nil {
		log.Println(err)
		utility.ResponseInternalServerError(c)
		return
	}
	if err != nil {
		log.Println(err)
		utility.ResponseInternalServerError(c)
		return
	}
	utility.Response(http.StatusOK, "OK", gin.H{
		"resume_id": data.Resumeid,
	}, c)
}

// 更新简历
func UpdateResume(c *gin.Context) {
	var data model.Resume
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utility.ResponseBadRequest(c)
		return
	}
	_, err = model.GetResumeByResumeidAndUserid(data.Resumeid, data.Userid)
	if err != nil {
		log.Println(err)
		utility.Response(404, "resume not found", nil, c)
		return
	}
	err = model.UpdateResume(&data)
	if err != nil {
		log.Println(err)
		utility.ResponseBadRequest(c)
		return
	}
	utility.Response(http.StatusOK, "OK", nil, c)
}

// 读取单个简历具体信息
func Read_resume(c *gin.Context) {
	userid := c.Query("user_id")
	resumeid := c.Query("id")
	userid_, _ := strconv.Atoi(userid)
	resumeid_, _ := strconv.Atoi(resumeid)
	resume, err := model.GetResumeByResumeidAndUserid(resumeid_, userid_)
	if err != nil {
		utility.Response(404, "not found resume", nil, c)
		return
	}
	utility.ResponseOK(c, gin.H{
		"user_id":              resume.Userid,
		"name":                 resume.Name,
		"phone":                resume.Phone,
		"user_email":           resume.Email,
		"graduate_data":        resume.Graduate_Data,
		"education_background": resume.Background,
		"previous_jobs":        resume.Previous_jobs,
		"work_experience":      resume.Experience,
		"self_introduce":       resume.Introduce,
		"personal_skills":      resume.Personal_skills,
		"resume_id":            resume.Resumeid,
		"creation_time":        resume.Date,
		"router":               resume.Router,
	})
}

// 读取所有简历
func Read_user(c *gin.Context) {
	userid := c.Query("user_id")
	userid_, _ := strconv.Atoi(userid)
	resumes, err := model.Getresumes(userid_)
	if err != nil {
		utility.ResponseBadRequest(c)
		return
	}
	list := make([]SummaryOFResumes, len(resumes))
	for i, value := range resumes {
		list[i].Title = value.Title
		list[i].Date = value.Date
		list[i].Resumeid = value.Resumeid
	}
	utility.ResponseOK(c, gin.H{
		"list":   list,
		"number": len(resumes),
	})
}
