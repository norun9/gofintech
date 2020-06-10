package main

import (
	"GoFintech/bitflyer"
	"GoFintech/config"
	"GoFintech/utils"
	"fmt"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	fmt.Println(apiClient.GetBalance())
}