FROM golang:1.20 AS builder

WORKDIR /app

COPY . .

RUN go env -w GOPROXY=https://proxy.golang.com.cn,direct
RUN go mod tidy
RUN go build -o myapp .

FROM alpine:latest

WORKDIR /app

# 从 builder 镜像中复制 Go 二进制文件到当前镜像中
COPY --from=builder /app/myapp .

# 定义容器启动时运行的命令
CMD ["./myapp"]
