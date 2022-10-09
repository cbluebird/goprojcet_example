package model

import (
	"resume/db"
)

type User struct {
	Userid   int    `json:"user_id"`
	UserName string `json:"user_name"`
	Email    string `json:"user_email"`
	Password string `json:"password"`
	Usersex  string `json:"user_sex"`
}

func GetUserByUsername(username string) (*User, error) {
	var user *User
	err := db.DB.Where("user_name = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	} else {
		return user, nil
	}
}

func GetUserByUserid(userid int) error {
	var user *User
	err := db.DB.Where("userid = ?", userid).First(&user).Error
	if err != nil {
		return err
	} else {
		return nil
	}
}

func AddUser(user *User) error {
	return db.DB.Create(user).Error
}

func UpdateUser(user *User) error {
	err := db.DB.Model(User{}).Where(
		User{
			Userid: user.Userid,
		}).Updates(&user).Error
	return err
}

func UpdateUserByEmail(user *User) error {
	err := db.DB.Model(User{}).Where(
		User{
			Email: user.Email,
		}).Updates(&user).Error
	return err
}
