package main

import (
	"GoFintech/app/models"
	"fmt"
	"GoFintech/config"
	"GoFintech/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	fmt.Println(models.DbConnection)
}