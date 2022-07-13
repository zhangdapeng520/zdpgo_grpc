package main

import (
	"context"
	"net"

	"github.com/zhangdapeng520/zdpgo_grpc/examples/helloworld/proto"

	"google.golang.org/grpc"
)

// 实现服务
type Server struct{}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{
		Message: "Hello " + request.Name,
	}, nil
}

func main() {
	// 创建服务
	g := grpc.NewServer()

	// 实例化服务
	proto.RegisterGreeterServer(g, &Server{})

	// 启动
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		panic(err)
	}
	err = g.Serve(lis)
	if err != nil {
		panic(err)
	}
}
