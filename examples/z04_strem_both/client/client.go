package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"google.golang.org/grpc"
	"zdpgo_grpc/examples/proto"
)

func main() {
	// 连接grpc
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// 创建客户端
	c := proto.NewStreamBothServerClient(conn)

	// 发送请求，获取响应
	response, _ := c.BothStream(context.Background())

	// 创建等待组
	wg := sync.WaitGroup{}
	wg.Add(2)

	// 不停的接收数据
	go func() {
		defer wg.Done()
		for {
			data, _ := response.Recv()
			fmt.Println("收到服务器信息：", data.Data)
		}
	}()

	// 不停的发送数据
	go func() {
		defer wg.Done()
		for {
			_ = response.Send(&proto.StreamBothRequest{Data: "我是客户端。。。"})
			time.Sleep(time.Second * 1)
		}
	}()

	wg.Wait()
}
