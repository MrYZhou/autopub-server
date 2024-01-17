打包
go build -ldflags '-w -s' .
upx压缩

暂时不推荐docker，因为相当于又套了一层，导致命令运行
需要增加目录映射，同时还必须在docker内部装环境。然后才可以
跑。等待后续优化

