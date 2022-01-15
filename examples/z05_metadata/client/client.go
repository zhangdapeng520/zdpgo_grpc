package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"zdpgo_grpc/examples/proto"
)

func main() {
	//stream
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)

	// 创建metadata
	md := metadata.New(map[string]string{
		"name":    "zhangdapeng",
		"pasword": "zhangdapeng",
	})

	// 将metadata注入到请求上下文对象中
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	// 携带metadata发送请求
	r, err := c.SayHello(ctx, &proto.HelloRequest{Name: "张大鹏"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}
