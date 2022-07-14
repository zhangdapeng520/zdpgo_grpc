package main

import (
	"fmt"
	"net"
	"time"

	"github.com/zhangdapeng520/zdpgo_grpc/examples/server_stream/proto"
	"google.golang.org/grpc"
)

// 服务对象
type server struct {
}

// 实现方法
func (s *server) GetStream(req *proto.StreamReqData, res proto.Greeter_GetStreamServer) error {
	i := 0
	for {
		i++
		_ = res.Send(&proto.StreamResData{
			Data: fmt.Sprintf("%v", time.Now().Unix()),
		})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}

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
