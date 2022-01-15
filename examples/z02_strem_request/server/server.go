package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
	"zdpgo_grpc/examples/proto"
)

type Server struct {
	proto.UnimplementedStreamRequestServerServer
}

// 实现接口方法，只需要有一个请求流服务器
func (s *Server) ClientStream(server proto.StreamRequestServer_ClientStreamServer) error {
	// 循环不断的接收客户端的数据
	for {
		if request, err := server.Recv(); err != nil {
			fmt.Println("获取客户端请求数据失败：", err)
			return err
		} else {
			fmt.Println("获取客户端请求数据成功：", request.Data)
		}
	}
}

func main() {
	// 创建服务
	g := grpc.NewServer()

	// 注册服务
	proto.RegisterStreamRequestServerServer(g, &Server{})

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
