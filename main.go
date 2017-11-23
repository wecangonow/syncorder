package main

import (
	"syncorder/config"

	"syncorder/httpserver"
	"syncorder/logger"
)

func init() {
	config.Init()
}

func main() {

	logger := logger.GetInstance()
	logger.Println("Start listening on port 8080")

	httpserver.Start()
}
