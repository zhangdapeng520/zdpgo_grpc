package main

import (
	"context"
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc"
	"zdpgo_grpc/examples/proto"
)

type Server struct {
	proto.UnsafeGreeterServer
}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply,
	error) {
	time.Sleep(2 * time.Second)
	return &proto.HelloReply{
		Message: "hello " + request.Name,
	}, nil
}

func main() {
	// 创建一个拦截器
	interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		// 请求之前的逻辑
		fmt.Println("接收到了一个新的请求")

		// 执行请求
		res, err := handler(ctx, req)

		// 请求之后的逻辑
		fmt.Println("请求已经完成")

		// 返回响应
		return res, err
	}

	// 使用拦截器
	opt := grpc.UnaryInterceptor(interceptor)

	// 创建服务，带拦截器
	g := grpc.NewServer(opt)

	// 注册服务
	proto.RegisterGreeterServer(g, &Server{})

	// 监听端口
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		panic("监听端口失败：" + err.Error())
	}

	// 启动服务
	err = g.Serve(lis)
	if err != nil {
		panic("启动服务失败：" + err.Error())
	}
}
