package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc"

	"github.com/zhangdapeng520/zdpgo_grpc/examples/metadata/proto"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply,
	error) {

	// 获取metadata对象
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		fmt.Println("get metadata error")
	}

	// 从metadata中提取值，每个值都是一个数组
	if nameSlice, ok := md["name"]; ok {
		fmt.Println(nameSlice)
		for i, e := range nameSlice {
			fmt.Println(i, e)
		}
	}

	// 返回响应
	return &proto.HelloReply{
		Message: "hello " + request.Name,
	}, nil
}

func main() {
	// 创建服务
	g := grpc.NewServer()

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
