package main

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"

	"github.com/zhangdapeng520/zdpgo_grpc/examples/client_stream/proto"
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

	// 客户端向服务器推送数据
	putS, _ := c.PutStream(context.Background())
	i := 0
	for {
		i++
		_ = putS.Send(&proto.StreamReqData{
			Data: fmt.Sprintf("客户端流数据 %d", i),
		})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}
}
