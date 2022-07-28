package main

import (
	"github.com/da-maltsev/UDV_Go_test/app"
)

func main() {

	app := &app.App{}
	app.Initialize()
	app.Run(":3001")
}
