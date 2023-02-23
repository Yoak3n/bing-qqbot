package main

import (
	"Bing-QQBot/config"
	"Bing-QQBot/router"
)

func init() {
	config.LoadConfig()

}

func main() {
	router.NewRouter()

}
