package zdpgo_grpc

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
)

/*
@Time : 2022/5/21 16:13
@Author : 张大鹏
@File : run.go
@Software: Goland2021.3.1
@Description: run类型的运行方法
*/

// RunServer 运行GRPC服务
func (g *Grpc) RunServer(
	registerFunc func(grpcServerRegister grpc.ServiceRegistrar, implGrpcServer interface{}),
	implServerArg interface{},
	interceptors ...grpc.ServerOption,
) {

	// 创建监听器
	address := fmt.Sprintf("%s:%d", g.Config.Server.Host, g.Config.Server.Port)
	listen, err := net.Listen("tcp", address)
	if err != nil {
		g.Log.Error("创建监听器失败", "error", err, "address", address)
		return
	}

	// 创建服务
	server := grpc.NewServer(interceptors...)

	// 注册服务
	registerFunc(server, implServerArg)
	g.Log.Debug("服务注册成功，即将启动GRPC服务", "address", address)

	// 启动服务
	if err = server.Serve(listen); err != nil {
		g.Log.Fatal("启动GRPC服务失败", "error", err, "address", address)
	}
}
