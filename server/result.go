package api

import (
	"github.com/gofiber/fiber/v2"
)

/**
说实话不是很满意,但是这是go不支持类名调用方法，反射鸡肋,
方法重载也不支持,没法写方法名相同,不同参数列表导致的
*/

type MyCtx struct {
	*fiber.Ctx
	Msg  string
	Code int
}

func AppResult(c *fiber.Ctx) *MyCtx {
	myCtx := c.Locals("myctx").(*MyCtx)
	return myCtx
}

func (c *MyCtx) Success(data interface{}, msg string) error {
	if msg == "" {
		msg = "success"
	}
	c.Msg = msg
	c.Code = 200
	return c.Response(data)
}

func (c *MyCtx) Fail(data interface{}, msg string) error {
	if msg == "" {
		msg = "fail"
	}
	c.Msg = msg
	c.Code = 400
	return c.Response(data)
}

func Success(c *fiber.Ctx, data interface{}) error {
	myCtx := c.Locals("myctx").(*MyCtx)
	myCtx.Msg = "success"
	myCtx.Code = 200
	return myCtx.Response(data)
}

func Fail(c *fiber.Ctx, data interface{}) error {
	myCtx := c.Locals("myctx").(*MyCtx)
	myCtx.Msg = "fail"
	myCtx.Code = 400
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
		"code": c.Code,
	}
	return c.JSON(res)
}
