package main

import (
	"context"
	"server/proto"

	"net"

	"github.com/zhangdapeng520/zdpgo_password"
	"google.golang.org/grpc"
)

type ServerAes struct {
	proto.UnimplementedServerAesServer
}

// Encrypt 对数据进行加密
func (s *ServerAes) Encrypt(ctx context.Context, request *proto.EncryptRequest) (*proto.EncryptResponse, error) {
	return &proto.EncryptResponse{Base64Encrypt: request.Base64Data}, nil
}

// EncryptString 对字符串进行加密
func (s *ServerAes) EncryptString(ctx context.Context, request *proto.EncryptStringRequest) (*proto.EncryptStringResponse, error) {
	aes := zdpgo_password.NewAes(zdpgo_password.AesConfig{})
	encryptData := aes.EncryptString(request.Data)
	return &proto.EncryptStringResponse{Base64Encrypt: encryptData}, nil
}

// DecryptString 对字符串进行解密
func (s *ServerAes) DecryptString(ctx context.Context, request *proto.DecryptStringRequest) (*proto.DecryptStringResponse, error) {
	aes := zdpgo_password.NewAes(zdpgo_password.AesConfig{})
	decryptData, _ := aes.DecryptString(request.Base64Encrypt)
	return &proto.DecryptStringResponse{DecryptData: decryptData}, nil
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
