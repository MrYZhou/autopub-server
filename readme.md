打包
go build -ldflags '-w -s' .
upx压缩

暂时不推荐docker，因为相当于又套了一层，导致命令运行
需要增加目录映射，同时还必须在docker内部装环境。然后才可以
跑。等待后续优化

开发文档

项目结构
router controler层写路由,前置校验，少部分简单的业务逻辑
server service层,写业务逻辑
test 测试
util 工具和系统结构和初始化相关代码

启动:
方式一：有下载air
air
方式二：
go run .
