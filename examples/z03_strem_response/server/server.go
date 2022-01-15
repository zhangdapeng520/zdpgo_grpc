package main

import (
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc"
	"zdpgo_grpc/examples/proto"
)

type Server struct{}

// 实现接口方法
// 参数1：客户端请求对象
// 参数2：服务端流响应服务对象
func (s *Server) ServerStream(request *proto.StreamResponseRequest, server proto.StreamResponseServer_ServerStreamServer) error {
	// 接收客户端请求
	switch request.Data {
	case "go":
		// 循环不断的向客户端发送数据
		// 发送10次“我爱Go语言！！！”
		for i := 0; i < 10; i++ {
			data := proto.StreamResponseResponse{
				Data: fmt.Sprintf("第%d次：我爱Go语言！！！", i+1),
			}
			_ = server.Send(&data)
			time.Sleep(time.Second * 1)
		}
	case "study":
		// 发送10次“我爱学习！！！”
		for i := 0; i < 10; i++ {
			data := proto.StreamResponseResponse{
				Data: fmt.Sprintf("第%d次：我爱学习！！！", i+1),
			}
			_ = server.Send(&data)
			time.Sleep(time.Second * 1)
		}
	}

	return nil
}

func main() {
	// 创建服务
	g := grpc.NewServer()

	// 注册服务
	proto.RegisterStreamResponseServerServer(g, &Server{})

	// 监听端口
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
