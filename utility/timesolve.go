package utility

import (
	"gorm.io/gorm/utils"
	"time"
)

func GetData() string {
	now := time.Now()
	year := utils.ToString(now.Year())
	month := now.Format("01")
	day := utils.ToString(now.Format("02"))
	return year + "_" + month + "_" + day
}
