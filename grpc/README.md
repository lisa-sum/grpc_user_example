## 安装

1. 安装 `proto`与`protobuf`
```shell
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
```

2. 安装`gRPC`
```shell
go get -u google.golang.org/grpc
```

3. 安装`gRPC`生成 `go` 语言的插件
```shell
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## 使用

示例:
使用 `gPRC` 来生成 `proto`,生成的文件在对应文件所在的目录
```shell
protoc -I=. --go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
*/*.proto
```
