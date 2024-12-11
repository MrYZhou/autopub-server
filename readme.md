## 项目结构

适合微小型项目, 10 个左右的业务模块<br />

router 写路由,前置校验,业务逻辑<br />
test 测试代码<br />
util 系统和初始化相关代码

## 项目环境

### 安装依赖

代理

```
go env -w GOPROXY=https://goproxy.cn,direct
```

```
go install
```

### 启动

```
 f5调试启动
```

<a name="ka6CV"></a>

## 部署

<a name="LsvlL"></a>

### 方式一.docker

```
docker build  -t fiber .
```

```
docker run  -d -p 8085:8083 --name fiber --restart always -e mode=production -e dbUrl="root:root@tcp(192.168.3.1:3306)/study" fiber
```

<a name="JAkhv"></a>

### 方式二.本地打包

```
go build -ldflags "-w -s" .
```

upx 压缩(可进一步压缩打包大小)

https://github.com/upx/upx/releases

```
upx autopub-server.exe
```
