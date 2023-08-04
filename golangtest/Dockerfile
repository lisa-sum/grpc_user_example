FROM golang:1.20

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
# 预先复制/缓存 go.mod用于预下载依赖项，并且只有在它们更改后的后续构建中才重新下载它们
COPY go.mod go.sum ./

# 源
RUN go env -w GOPROXY=https://goproxy.cn,direct

# 下载依赖并验证
RUN go mod download && go mod verify

COPY . .

# 编译
# -v 详细输出编译结果
# -o 指定输出文件名
# ./... 递归编译当前目录下所有包
RUN go build -v -o /app/web ./...

CMD ["/app/web"]

# 打包镜像
# docker build -t goapp .

# 只运行一次镜像, 退出后删除
# docker run -it --rm --name goapp -p 8080:8080 goapp

# 后台运行镜像
# docker run -d --name goapp -p 8080:8080 goapp
