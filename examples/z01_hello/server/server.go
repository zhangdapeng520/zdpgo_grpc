package main

import (
	"context"
	"log"
	"net"

	pb "zdpgo_grpc/examples/proto"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server 用于实现 helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello 实现 helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("接收到数据： %v", in.GetName())
	return &pb.HelloReply{Message: "你好 " + in.GetName()}, nil
}

func main() {
	// 创建监听器
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("监听服务失败： %v", err)
	}

	// 创建服务
	s := grpc.NewServer()

	// 注册服务
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("服务监听地址: %v", lis.Addr())

	// 启动服务
	if err := s.Serve(lis); err != nil {
		log.Fatalf("启动服务失败: %v", err)
	}
}
