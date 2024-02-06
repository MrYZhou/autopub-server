# 第一阶段：构建应用
FROM golang:bullseye
# 设置容器工作目录在 /app
WORKDIR /app
# 将本地的代码文件复制到容器中的工作目录, COPY . ./app 可以用点替代工作目录的意思
COPY . .
# 设置代理地址来加快下载依赖的速度
RUN go env -w GOPROXY=https://goproxy.cn,direct
# 下载依赖
RUN go mod download
# 执行命令
CMD ["go", "run", "main.go"]
# 暴露端口
EXPOSE 8083