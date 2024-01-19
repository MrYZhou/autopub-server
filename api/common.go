package api

import "github.com/gofiber/fiber/v2"

type AppResult struct {
	Message string
}


func (app *AppResult) Success(c *fiber.Ctx, model interface{}) error{
	app.Message ="success"
	return c.JSON(model)
}

func (app *AppResult) Fail(c *fiber.Ctx, model interface{}) error{
	app.Message ="fail"
	return c.JSON(model)
}
