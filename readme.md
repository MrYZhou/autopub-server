## 项目环境

基于 fiber 框架的微小型项目, mysql 数据库，gplus<br />

### 配置代理

```
go env -w GOPROXY=https://goproxy.cn,direct
```

### 安装依赖

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

upx --best --lzma autopub-server.exe
```
