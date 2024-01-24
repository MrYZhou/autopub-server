package api

import (
	"reflect"

	"github.com/gofiber/fiber/v2"
)

/**
痛点:
说实话不是很满意,但是这是go不支持类名调用方法，反射鸡肋,
方法重载也不支持,没法写方法名相同,不同参数列表
而且还没办法设置默认参数名,导致还要手动传一个值进来。
可选参数也没办法设置

解决:
最后终于找的使用不定参数的方式去简化我们的统一响应模型
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

func (c *MyCtx) Success(params ...interface{}) error {
	c.Code = 200
	c.Msg = "success"
	return c.judge(params)
}

func (c *MyCtx) Fail(params ...interface{}) error {
	c.Code = 400
	c.Msg = "fail"
	return c.judge(params)
}

func CtxMiddleware(c *fiber.Ctx) error {
	ctx := &MyCtx{Ctx: c}
	c.Locals("myctx", ctx)
	return c.Next()
}

func (c *MyCtx) Response(data interface{}) error {
	res := map[string]interface{}{
		"msg":  c.Msg,
		"code": c.Code,
	}
	if data != nil {
		res["data"] = data
	}
	return c.JSON(res)
}
func (c *MyCtx) judge(params ...interface{}) error {
	// 参数为一个的时候,判断是字符串还是对象
	if len(params) == 1 {
		paramType := reflect.TypeOf(params[0])
		if paramType.Kind() == reflect.String {
			c.Msg = params[0].(string)
			return c.Response(nil)
		}
		return c.Response(params[0])
	}
	// 参数为两个的时候
	if len(params) == 2 {
		data := params[0]
		msg := params[1].(string)
		if msg != "" {
			c.Msg = msg
		}
		return c.Response(data)
	}

	return c.Response(nil)
}