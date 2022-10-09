package main

import (
	"resume/config"
	"resume/db"
	"resume/db/model"
	"resume/router"
)

func main() {
	config.InitConfig()
	db.InitDB()
	model.InitModel()
	router.InitRouter()
	router.Router.Run(":" + config.Config.Server.Port)
}
