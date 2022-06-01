package main

import (
	"short-url/model"
	"short-url/server"
)

func main() {
	model.Setup()
	server.SetupAndListen()
}
