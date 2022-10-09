package router

import (
	"resume/controller"
)

func SetRouter() {
	Router.POST("user/register", controller.Register)
	Router.POST("user/login", controller.Login)
	Router.POST("user/information", controller.UpdateUser)
	Router.GET("resume/read", controller.Read_resume)
	Router.POST("user/add", controller.Add)
	Router.GET("user/read", controller.Read_user)
	Router.POST("resume/template", controller.UpdateResume)
	Router.POST("user/repassword", controller.UpdateUserByEmail)
	Router.POST("user/email", controller.EmailVerfication)
}
