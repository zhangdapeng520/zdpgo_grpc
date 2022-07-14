package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	"github.com/zhangdapeng520/zdpgo_grpc/examples/server_stream/proto"
)

func main() {
	// 连接服务
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// 创建客户端
	c := proto.NewGreeterClient(conn)

	// 从服务器拉取数据流
	res, _ := c.GetStream(context.Background(), &proto.StreamReqData{Data: "客户端"})
	for {
		a, err := res.Recv() //如果大家懂socket编程的话就明白 send recv
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(a.Data)
	}
}
