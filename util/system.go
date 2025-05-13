package util

import (
	"autopub-server/common"
	"log"
	"os"

	"github.com/bytedance/sonic"
	"github.com/gofiber/contrib/monitor"
	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/static"
)

var app *fiber.App

func App() *fiber.App {
	if app == nil {
		// 加载根目录下的.env文件
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		app = fiber.New(fiber.Config{
			JSONEncoder:       sonic.Marshal,
			JSONDecoder:       sonic.Unmarshal,
			ReduceMemoryUsage: true,
		})
		// 启用 CORS 中间件
		app.Use(cors.New(cors.Config{
			AllowOrigins:        []string{"*"},                                                        // 明确允许的域名列表
			AllowMethods:        []string{"GET", "POST", "HEAD", "PUT", "DELETE", "PATCH", "OPTIONS"}, // 覆盖所有常用方法
			AllowHeaders:        []string{"Origin", "Content-Type", "Accept", "Authorization"},        // 包含常见请求头
			AllowCredentials:    false,                                                                // 如需携带 Cookie 或认证头需启用
			AllowPrivateNetwork: true,
		}))
		// 注册自定义中间件以转换上下文
		app.Use(CtxMiddleware)
		// 静态文件服务
		app.Get("/file*", static.New(os.Getenv("resources")))
		// app.Get("/*", static.New(os.Getenv("resources")+"/dist")) // 兜底匹配其他路径
		// 监控
		app.Use("/metrics", monitor.New())
		err = common.OpenBrowser("http://" + os.Getenv("host"))
		if err != nil {
			log.Fatalf("无法打开浏览器: %v", err)
		}
	}
	return app
}
