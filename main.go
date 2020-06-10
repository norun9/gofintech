package main

import (
	"GoFintech/bitflyer"
	"GoFintech/config"
	"GoFintech/utils"
	"fmt"
	//"time"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)

	//order := &bitflyer.Order{
	//	ProductCode: config.Config.ProductCode,
	//	ChildOrderType: "MARKET",
	//	Side: "BUY",
	//	Size: 0.01,
	//	MinuteToExpires: 1,
	//	TimeInForce: "GTC",
	//}
	//res, _ := apiClient.SendOrder(order)
	//fmt.Println(res.ChildOrderAcceptanceID)

	i := "JRF20201012-144016-14058"
	params := map[string]string{
		"product_code": config.Config.ProductCode,
		"child_order_acceptance_id": i,
	}
	r, _ := apiClient.ListOrder(params)
	fmt.Println(r)
}