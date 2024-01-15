# 第一阶段：构建应用
FROM golang:bullseye AS builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go mod download
COPY . .
RUN  go build -o main .

# 第二阶段：制作干净的生产镜像
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]
EXPOSE 8083