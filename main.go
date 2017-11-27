package main

import (
	"syncorder/config"
	"syncorder/httpserver"
	"syncorder/task"
)

func init() {
	config.Init()
}

func main() {

	task.Start()
	httpserver.Start()

}
