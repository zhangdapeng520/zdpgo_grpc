package main

import (
	"context"
	"fmt"
	"time"

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
	c := proto.NewStreamRequestServerClient(conn)

	// 发送请求，获取响应
	client, _ := c.ClientStream(context.Background())

	// 发送10次 我爱Go语言
	for i := 0; i < 10; i++ {
		data := proto.StreamRequestRequest{
			Data: fmt.Sprintf("第%d次：我爱Go语言！！！", i+1),
		}
		client.Send(&data)
		time.Sleep(time.Second * 1)
	}
}
