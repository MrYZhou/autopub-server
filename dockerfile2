# 第一阶段：构建应用
FROM golang:bullseye AS builder
# 设置容器工作目录在 /app
WORKDIR /app
# 将本地的代码文件复制到容器中的工作目录, COPY . ./app 可以用点替代工作目录的意思
COPY . .
# 设置代理地址来加快下载依赖的速度
RUN go env -w GOPROXY=https://goproxy.cn,direct
# 下载依赖
RUN go mod download
# 编译应用
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -ldflags '-w -s' -o main .

# 第二阶段：制作干净的生产镜像
FROM alpine:latest
# https依赖
RUN apk --no-cache add ca-certificates
WORKDIR /app
# 从builder阶段产物中获取main可执行程序
COPY --from=builder /app/main .
# 执行命令
CMD ["./main"]
# 暴露端口
EXPOSE 8083