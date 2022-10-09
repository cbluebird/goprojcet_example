package model

import "resume/db"

type Resume struct {
	Userid          int    `json:"user_id"`
	Resumeid        int    `json:"resume_id"`
	Title           string `json:"title"`
	Date            string `json:"creation_time"`
	Name            string `json:"name"`
	Phone           string `json:"phone"`
	Email           string `json:"user_email"`
	Graduate_Data   string `json:"graduate_data"`
	Background      string `json:"education_background"`
	Previous_jobs   string `json:"previous_jobs"`
	Experience      string `json:"work_experience"`
	Introduce       string `json:"self_introduce"`
	Personal_skills string `json:"personal_skills"`
	Router          string `json:"router"`
}

func AddResume(resume *Resume) error {
	return db.DB.Create(resume).Error
}

func UpdateResume(resume *Resume) error {
	err := db.DB.Model(Resume{}).Where(
		Resume{
			Userid:   resume.Userid,
			Resumeid: resume.Resumeid,
		}).Updates(&resume).Error
	return err
}

// 通过简历名称和用户id查询
func GetResumeByResume_nameAndUserid(resume_name string, userid int) error {
	var user *Resume
	err := db.DB.Where("title = ? AND userid = ?", resume_name, userid).First(&user).Error
	if err != nil {
		return err
	} else {
		return nil
	}
}

func GetResumeByResumeidAndUserid(resumeid int, userid int) (*Resume, error) {
	var resume *Resume
	err := db.DB.Where("resumeid = ? AND userid = ?", resumeid, userid).First(&resume).Error
	if err != nil {
		return nil, err
	} else {
		return resume, nil
	}
}

func Getresumes(userid int) ([]Resume, error) {
	var resumes []Resume
	result := db.DB.Where(
		Resume{
			Userid: userid,
		}).Find(&resumes)
	if result.Error != nil {
		return nil, result.Error
	}

	return resumes, nil
}
