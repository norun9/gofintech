package main

import (
	"GoFintech/config"
	"GoFintech/utils"
	"log"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	log.Println("test")
}