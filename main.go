package main

import (
	"github.com/da-maltsev/UDV_Go_test/app"
	"github.com/da-maltsev/UDV_Go_test/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
