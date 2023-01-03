package main

import (
	"estudo-go/configs"
)

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}
