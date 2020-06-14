package main

import (
	"GoFintech/app/controllers"
	"GoFintech/config"
	"GoFintech/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	controllers.StreamIngestionData()
	controllers.StartWebServer()
}