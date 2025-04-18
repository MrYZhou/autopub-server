package util

import (
	"autopub-server/common"
	"log"
	"os"

	"github.com/bytedance/sonic"
	"github.com/gofiber/contrib/monitor"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
)

var app *fiber.App

func App() *fiber.App {
	if app == nil {
		app = fiber.New(fiber.Config{
			JSONEncoder: sonic.Marshal,
			JSONDecoder: sonic.Unmarshal,
		})
		// 注册自定义中间件以转换上下文
		app.Use(CtxMiddleware)
		// 静态文件服务
		app.Get("/file*", static.New(os.Getenv("resources")))
		app.Get("/*", static.New(os.Getenv("resources")+"/dist")) // 兜底匹配其他路径
		// 监控
		app.Use("/metrics", monitor.New())
		err := common.OpenBrowser("http://" + os.Getenv("host"))
		if err != nil {
			log.Fatalf("无法打开浏览器: %v", err)
		}
	}
	return app
}
