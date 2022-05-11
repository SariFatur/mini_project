package main

import (
	"myproject/config"
	"myproject/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Start(":80")
}
