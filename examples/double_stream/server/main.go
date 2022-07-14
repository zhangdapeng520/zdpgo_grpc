package main

import (
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/zhangdapeng520/zdpgo_grpc/examples/double_stream/proto"

	"google.golang.org/grpc"
)

// 服务
type server struct {
}

// 实现方法
func (s *server) AllStream(allStr proto.Greeter_AllStreamServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)

	// 不停的接收数据
	go func() {
		defer wg.Done()
		for {
			data, _ := allStr.Recv()
			fmt.Println("收到客户端消息：" + data.Data)
		}
	}()

	// 不停的发送数据
	go func() {
		defer wg.Done()
		for {
			_ = allStr.Send(&proto.StreamResData{Data: "我是服务器"})
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()
	return nil
}

func main() {
	// 创建监听
	lis, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		panic(err)
	}

	// 创建服务
	s := grpc.NewServer()

	// 注册服务
	proto.RegisterGreeterServer(s, &server{})

	// 启动服务
	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}
