package server

import (
	pb "basic/proto"
	"context"
	"log"
)

/*
@Time : 2022/5/21 16:30
@Author : 张大鹏
@File : server.go
@Software: Goland2021.3.1
@Description: grpc服务实现核心
*/

// Server 用于实现 helloworld.GreeterServer.
type Server struct {
	pb.UnimplementedGreeterServer
}

// SayHello 实现 helloworld.GreeterServer
func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("接收到数据： %v", in.GetName())
	return &pb.HelloReply{Message: "你好 " + in.GetName()}, nil
}
