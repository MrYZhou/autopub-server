
## 项目结构（适合小型项目10个左右的业务模块）
router 写路由,前置校验，少部分简单的业务逻辑<br />server 写业务逻辑<br />test 测试代码<br />util 系统和初始化相关代码

## 项目环境
### 安装依赖
代理
```
go env -w GOPROXY=https://goproxy.cn,direct
```
```
go get 或 go install
```
### 启动
方式一：下载air([https://github.com/cosmtrek/air)](https://github.com/cosmtrek/air))，热部署启动<br />

在外面cmd窗口执行
```
go install github.com/air-verse/air@latest
```

项目执行 
```
air init
```
项目启动 
```
air
```

方式二：
```
go run .
```

方式三: f5调试启动

<a name="ka6CV"></a>
## 部署
<a name="LsvlL"></a>

### 方式一.docker
```
docker build -f dockerfile -t fiber .
```
```
docker run  -d -p 8085:8083 --name fiber --restart always -e mode=production -e dbUrl="root:root@tcp(192.168.3.1:3306)/study" fiber
```

<a name="JAkhv"></a>

### 方式二.本地打包
```
go build -ldflags "-w -s" .
```
upx压缩(可进一步压缩打包大小)

https://github.com/upx/upx/releases
```
upx autopub-server.exe
```