package api

import (
	"github.com/gofiber/fiber/v2"
)

type MyCtx struct {
	*fiber.Ctx
	Msg string
}

func Success(c *fiber.Ctx, data interface{}) error {
	myCtx := c.Locals("myctx").(*MyCtx)
	myCtx.Msg = "success"
	return myCtx.Response(data)
}

func Fail(c *fiber.Ctx, data interface{}) error {
	myCtx := c.Locals("myctx").(*MyCtx)
	myCtx.Msg = "fail"
	return myCtx.Response(data)
}

func CtxMiddleware(c *fiber.Ctx) error {
	ctx := &MyCtx{Ctx: c}
	c.Locals("myctx", ctx)
	return c.Next()
}

func (c *MyCtx) Response(data interface{}) error {
	res := map[string]interface{}{
		"msg":  c.Msg,
		"data": data,
	}
	return c.JSON(res)
}
