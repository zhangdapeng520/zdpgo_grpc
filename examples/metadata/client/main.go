package main

import (
	"context"
	"fmt"

	"github.com/zhangdapeng520/zdpgo_grpc/examples/metadata/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	// 连接服务
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// 创建客户端
	c := proto.NewGreeterClient(conn)

	// 创建metdata
	//md := metadata.Pairs("timestamp", time.Now().Format(timestampFormat))
	md := metadata.New(map[string]string{
		"name":    "zhangdapeng520",
		"pasword": "zhangdapeng520",
	})

	// 将metdata注入到context
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	// 携带metadata发送请求
	r, err := c.SayHello(ctx, &proto.HelloRequest{Name: "张大鹏"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}
