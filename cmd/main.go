package main

import (
	"gin-mall/config"
	"gin-mall/routers"
)

func main() {
	config.Init()
	r := routers.NewRouter()
	_ = r.Run(config.HttpPort)
}
