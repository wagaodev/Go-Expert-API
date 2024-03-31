package main

import (
	"github.com/wagaodev/Go-Expert-API/configs"
)

func main() {
	config, _ := configs.Init(".")
	println(config.DBDriver)
}
