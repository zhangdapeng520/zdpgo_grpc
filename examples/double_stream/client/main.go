package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"google.golang.org/grpc"

	"github.com/zhangdapeng520/zdpgo_grpc/examples/double_stream/proto"
)

func main() {
	// 创建连接
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// 创建客户端
	c := proto.NewGreeterClient(conn)

	// 双向流模式
	allStr, _ := c.AllStream(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(2)

	// 不停的接收数据
	go func() {
		defer wg.Done()
		for {
			data, _ := allStr.Recv()
			fmt.Println("收到客户端消息：" + data.Data)
		}
	}()

	// 不停的发送数据
	go func() {
		defer wg.Done()
		for {
			_ = allStr.Send(&proto.StreamReqData{Data: "客户端数据流"})
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
}
