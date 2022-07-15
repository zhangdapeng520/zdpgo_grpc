package main

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"

	"github.com/zhangdapeng520/zdpgo_grpc/examples/interpretor/proto"
)

func main() {
	// 创建客户端的拦截器
	interceptor := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		start := time.Now()

		// 执行请求
		err := invoker(ctx, method, req, reply, cc, opts...)

		// 请求之后的逻辑
		fmt.Printf("耗时：%s\n", time.Since(start))
		return err
	}

	// 创建拦截器列表
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithUnaryInterceptor(interceptor))

	// 创建连接
	conn, err := grpc.Dial("127.0.0.1:50051", opts...)
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
