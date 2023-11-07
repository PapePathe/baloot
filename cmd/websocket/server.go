package main

import (
	"pathe.co/zinx/znet"
)

func main() {
	s := znet.NewSocketHandler()
	app := znet.NewSocketApp("7777")

	app.SetupRoutes(s)
	app.Start()
}
