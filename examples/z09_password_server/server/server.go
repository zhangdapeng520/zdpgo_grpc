package main

import (
	"context"
	"server/proto"

	"net"

	"google.golang.org/grpc"
)

type ServerAes struct {
	proto.UnimplementedServerAesServer
}

func (s *ServerAes) Encrypt(ctx context.Context, request *proto.EncryptRequest) (*proto.EncryptResponse, error) {
	return &proto.EncryptResponse{Base64Encrypt: request.Base64Data}, nil
}

func main() {
	// 创建服务
	g := grpc.NewServer()

	// 注册服务
	proto.RegisterServerAesServer(g, &ServerAes{})

	// 监听服务
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
