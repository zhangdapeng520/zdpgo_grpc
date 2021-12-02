package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"zgo_grpc/examples/proto"
)

func main() {
	// 连接grpc
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// 创建客户端
	c := proto.NewGreeterClient(conn)

	// 发送请求，获取响应
	r, err := c.SayHello(context.Background(), &proto.HelloRequest{Name: "zhangdapeng"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}