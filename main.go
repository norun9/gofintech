package main

import (
	"GoFintech/config"
	"fmt"
)

func main() {
	fmt.Println(config.Config.ApiKey)
	fmt.Println(config.Config.ApiSecret)
}