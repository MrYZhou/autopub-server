<a name="YVSBY"></a>
## 项目结构
router controler层写路由,前置校验，少部分简单的业务逻辑<br />server service层,写业务逻辑<br />test 测试代码<br />util 系统和初始化相关代码
<a name="hhs9E"></a>
## 项目环境
<a name="tqDxu"></a>
### 安装依赖
go get 或 go install
<a name="c7UlQ"></a>
### 启动
方式一：下载air([https://github.com/cosmtrek/air)](https://github.com/cosmtrek/air))，热部署启动<br />

go install github.com/cosmtrek/air@latest <br />

air<br />方式二：<br />go run .

<a name="ka6CV"></a>
## 部署
<a name="fd8wE"></a>
### 构建测试镜像
docker build -f dockerfile-test -t fibertest .
<a name="LsvlL"></a>
### 构建正式镜像
docker build -f dockerfile -t fiber .

### 运行镜像
docker run  -d -p 8085:8083 --name fiber --restart always 
-e mode=production -e dbUrl="root:root@tcp(192.168.3.1:3306)/study" 
-e GOPROXY=https://goproxy.cn,direct fiber

<a name="JAkhv"></a>
### 本地打包
go build -ldflags '-w -s' .<br />upx压缩(可选进一步压缩打包大小)
