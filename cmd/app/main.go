package main

import (
	"estudo-go/configs"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	println(config.DBDriver)
}
