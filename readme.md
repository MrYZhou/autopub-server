打包
go build -ldflags '-w -s' .
upx压缩(可选进一步压缩打包大小)


项目结构
router controler层写路由,前置校验，少部分简单的业务逻辑
server service层,写业务逻辑
test 测试
util 工具和系统结构和初始化相关代码

安装包
go get 或 go install

启动:
方式一：有下载air(https://github.com/cosmtrek/air)
air
方式二：
go run .

构建测试镜像
docker build -f dockerfile-test -t fibertest .
构建正式镜像
docker build -f dockerfile -t fiber .
