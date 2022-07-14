package main

import (
	"fmt"
	"net"

	"github.com/zhangdapeng520/zdpgo_grpc/examples/client_stream/proto"
	"google.golang.org/grpc"
)

// 服务对象
type server struct {
}

// 实现服务方法
func (s *server) PutStream(cliStr proto.Greeter_PutStreamServer) error {
	for {
		if a, err := cliStr.Recv(); err != nil {
			fmt.Println(err)
			break
		} else {
			fmt.Println(a.Data)
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
