package main

import (
	"tech/internal/app"
	"tech/internal/entities/utils"
)

func main() {
	utils.InitMutex()
	app.RunApp()
}
