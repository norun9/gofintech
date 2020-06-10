package main

import (
	"GoFintech/bitflyer"
	"GoFintech/config"
	"GoFintech/utils"
	"fmt"
	"time"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	ticker, _ := apiClient.GetTicker("BTC_USD")
	fmt.Println(ticker)
	fmt.Println(ticker.GetMidPrice())
	fmt.Println(ticker.TruncateDateTime(time.Hour))
}