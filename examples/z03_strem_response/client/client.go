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
	c := proto.NewStreamResponseServerClient(conn)

	// 发送请求，获取响应
	request := proto.StreamResponseRequest{
		Data: "go",
	}
	response, _ := c.ServerStream(context.Background(), &request)

	// 循环不断的接收服务端的数据
	for {
		data, err := response.Recv()
		if err != nil {
			fmt.Println("接收服务端数据失败：", err)
			break
		} else {
			fmt.Println("接收服务端数据成功：", data)
		}
	}

	// 修改参数
	request.Data = "study"
	response, _ = c.ServerStream(context.Background(), &request)

	// 循环不断的接收服务端的数据
	for {
		data, err := response.Recv()
		if err != nil {
			fmt.Println("接收服务端数据失败：", err)
			break
		} else {
			fmt.Println("接收服务端数据成功：", data)
		}
	}
}
