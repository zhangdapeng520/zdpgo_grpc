package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"

	"github.com/zhangdapeng520/zdpgo_grpc/examples/interpretor/proto"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply,
	error) {
	return &proto.HelloReply{
		Message: "hello " + request.Name,
	}, nil
}

func main() {
	// 创建拦截器
	interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		// 执行之前的逻辑
		fmt.Println("接收到了一个新的请求")

		// 放行请求
		res, err := handler(ctx, req)

		// 执行之后的逻辑
		fmt.Println("请求已经完成")

		// 返回
		return res, err
	}

	// 创建拦截器
	opt := grpc.UnaryInterceptor(interceptor)

	// 创建服务
	g := grpc.NewServer(opt)

	// 注册服务
	proto.RegisterGreeterServer(g, &Server{})

	// 创建监听
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		panic("failed to listen:" + err.Error())
	}

	// 启动服务
	err = g.Serve(lis)
	if err != nil {
		panic("failed to start grpc:" + err.Error())
	}
}
