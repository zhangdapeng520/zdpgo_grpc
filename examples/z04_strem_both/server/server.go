package main

import (
	"fmt"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
	"zgo_grpc/examples/proto"
)

type Server struct{}

// 实现接口方法
// 参数1：双向数据流服务对象（rpc接口方法）
func (s *Server) BothStream(server proto.StreamBothServer_BothStreamServer) error {
	// 创建等待组
	wg := sync.WaitGroup{}
	wg.Add(2)

	// 不断的读数据
	go func() {
		defer wg.Done()
		for {
			data, _ := server.Recv()
			fmt.Println("接收到客户端消息：", data.Data)
		}
	}()

	// 不断的写数据
	go func() {
		defer wg.Done()
		for {
			_ = server.Send(&proto.StreamBothResponse{
				Data: "来自服务器的数据。。。",
			})
			time.Sleep(time.Second * 1)
		}
	}()

	// 等待组中go程完成
	wg.Wait()
	return nil
}

func main() {
	// 创建服务
	g := grpc.NewServer()

	// 注册服务
	proto.RegisterStreamBothServerServer(g, &Server{})

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
