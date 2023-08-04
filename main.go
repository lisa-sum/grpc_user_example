package main

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"log"
	"net"
	"protobuf/proto/user"
)

// IServer server接口
type IServer interface {
	mustEmbedUnimplementeduserServiceServer()
	Login(ctx context.Context, req *user.LoginRequest) (*user.LoginResponse, error)
	Register(ctx context.Context, req *user.RegisterRequest) (*user.RegisterResponse, error)
}

type server struct {
	user.UnimplementedUserServiceServer
}

func (s server) mustEmbedUnimplementeduserServiceServer() {}

func (s server) Login(ctx context.Context, req *user.LoginRequest) (*user.LoginResponse, error) {
	if req.Username == "" && req.Password == "" {
		return nil, errors.New("非空账号密码!")
	}
	return &user.LoginResponse{
		StateCode: 200,
		StateMsg:  "Ok",
	}, nil
}

// Register 实现了注册操作
func (s server) Register(ctx context.Context, req *user.RegisterRequest) (*user.RegisterResponse, error) {
	// 模拟注册成功
	return &user.RegisterResponse{
		StateCode: 200,
		StateMsg:  "OK",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":4000")
	if err != nil {
		log.Fatal(err)
	}

	gs := grpc.NewServer()

	user.RegisterUserServiceServer(gs, &server{})

	err = gs.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
