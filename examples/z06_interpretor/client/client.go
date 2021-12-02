package main

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"zgo_grpc/examples/proto"
)

func main() {
	// 创建客户端的拦截器
	interceptor := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		// 请求之前的逻辑
		start := time.Now()

		// 执行请求
		err := invoker(ctx, method, req, reply, cc, opts...)

		// 请求之后的逻辑
		fmt.Printf("耗时：%s\n", time.Since(start))

		// 返回请求结果
		return err
	}

	// 请求要携带的参数
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithUnaryInterceptor(interceptor))
	conn, err := grpc.Dial("127.0.0.1:50051", opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// 创建客户端并发送请求
	c := proto.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &proto.HelloRequest{Name: "zhangdapeng"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}
