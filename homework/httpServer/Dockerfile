FROM golang:alpine AS builder

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
	GOPROXY=https://goproxy.cn

# 移动到工作目录：/build
WORKDIR /build

# 将我们的代码编译成二进制可执行文件 app
# 将代码复制到容器中
COPY . .
RUN go build -o src/main.go httpServer

###################
# 接下来创建一个小镜像
###################
FROM scratch

# 从builder镜像中把/dist/app 拷贝到当前目录
COPY --from=builder /build/httpServer /

#外放端口
EXPOSE 8001

# 需要运行的命令
ENTRYPOINT ["./httpServer"]

