package util

import "github.com/gofiber/fiber/v2"

func Init() {
	DbInit()
}

var app  *fiber.App

func App() *fiber.App{
	if app == nil{
		app = fiber.New()
	}
	return app
}