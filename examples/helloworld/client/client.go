package main

import (
	"context"
	"fmt"

	"github.com/zhangdapeng520/zdpgo_grpc/examples/helloworld/proto"
	"google.golang.org/grpc"
)

func main() {
	// 导入grpc

	// 连接
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// 创建客户端
	c := proto.NewGreeterClient(conn)

	// 调用
	r, err := c.SayHello(context.Background(), &proto.HelloRequest{Name: "张大鹏"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}
