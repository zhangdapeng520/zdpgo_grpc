package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"zgo_grpc/examples/proto"

	"time"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)

	// 设置超时
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)

	// 发送请求获取数据
	_, err = c.SayHello(ctx, &proto.HelloRequest{Name: "zhangdapeng"})

	// 处理错误
	if err != nil {

		// 解析错误
		st, ok := status.FromError(err)
		if !ok {
			panic("解析error失败")
		}

		// 提取错误信息和错误码
		fmt.Println(st.Message())
		fmt.Println(st.Code())
	}
}
