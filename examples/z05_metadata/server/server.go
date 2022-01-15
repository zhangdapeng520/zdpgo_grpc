package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc/metadata"
	"zdpgo_grpc/examples/proto"

	"net"

	"google.golang.org/grpc"
)

type Server struct {
	proto.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply,
	error) {
	// 获取metadata中的数据
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		fmt.Println("获取metadata失败")
	}

	// 查看metadata中都有什么
	for k, v := range md {
		fmt.Println("metadata中的数据：", k, v)
	}

	// 获取单个值，得到的是一个切片
	if nameSlice, ok := md["name"]; ok {
		fmt.Println(nameSlice)
		for i, e := range nameSlice {
			fmt.Println(i, e)
		}
	}

	// 正常的业务逻辑
	return &proto.HelloReply{
		Message: "hello " + request.Name,
	}, nil
}

func main() {
	// 创建服务
	g := grpc.NewServer()

	// 注册服务
	proto.RegisterGreeterServer(g, &Server{})

	// 监听端口
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		panic("监听端口失败:" + err.Error())
	}

	// 启动服务
	err = g.Serve(lis)
	if err != nil {
		panic("启动grpc服务失败:" + err.Error())
	}
}
