package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"resume/db"
	"resume/db/model"
	"resume/utility"
)

type LoginData struct {
	Username string `json:"user_name"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var data LoginData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		utility.ResponseBadRequest(c)
		return
	}
	var user *model.User
	user, err = model.GetUserByUsername(data.Username)
	if user == nil {
		utility.Response(http.StatusNotFound, "User not found", nil, c)
		return
	}
	if !utility.PasswordVerify(data.Password, user.Password) {
		utility.Response(http.StatusBadRequest, "Wrong Password", nil, c)
		return
	}
	fmt.Println("Login: ", user.Userid)
	utility.Response(http.StatusOK, "OK", gin.H{"user_id": user.Userid}, c)
}

func Register(c *gin.Context) {
	var data model.User
	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Println(err)
		utility.ResponseBadRequest(c)
		return
	}
	_, err = model.GetUserByUsername(data.UserName)
	if err == nil {
		log.Println(err)
		utility.Response(404, "username repeat", nil, c)
		return
	}
	data.Password, err = utility.PasswordHash(data.Password)
	if err != nil {
		log.Println(err)
		utility.ResponseInternalServerError(c)
		return
	}
	var last model.User
	db.DB.Last(&last)
	data.Userid = last.Userid + 1
	err = model.AddUser(&data)
	if err != nil {
		log.Println(err)
		utility.ResponseInternalServerError(c)
		return
	}
	utility.Response(http.StatusOK, "OK", nil, c)
}

func UpdateUser(c *gin.Context) {
	var data model.User
	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Println(err)
		utility.ResponseInternalServerError(c)
		return
	}
	data.Password, err = utility.PasswordHash(data.Password)
	if err != nil {
		log.Println(err)
		utility.ResponseInternalServerError(c)
		return
	}
	err = model.UpdateUser(&data)
	if err != nil {
		log.Println(err)
		utility.ResponseInternalServerError(c)
		return
	}
	utility.Response(http.StatusOK, "OK", nil, c)
}

func UpdateUserByEmail(c *gin.Context) {
	var data model.User
	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Println(err)
		utility.ResponseInternalServerError(c)
		return
	}
	data.Password, err = utility.PasswordHash(data.Password)
	if err != nil {
		log.Println(err)
		utility.ResponseInternalServerError(c)
		return
	}
	err = model.UpdateUserByEmail(&data)
	if err != nil {
		log.Println(err)
		utility.Response(404, "email not found", nil, c)
		return
	}
	utility.Response(http.StatusOK, "OK", nil, c)
}
