package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"protobuf/proto/user"
)

func main() {
	conn, err := grpc.Dial("localhost:4000", grpc.WithInsecure())
	defer conn.Close()

	client := user.NewUserServiceClient(conn)

	result, err := client.Login(context.Background(), &user.LoginRequest{
		Username: "123",
		Password: "23",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result)
}
