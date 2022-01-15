package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"zdpgo_grpc/examples/proto"

	"net"

	"google.golang.org/grpc"
)

type Server struct {
	proto.UnsafeGreeterServer
}

// 实现服务方法
func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply,
	error) {
	return &proto.HelloReply{
		Message: "hello " + request.Name,
	}, nil
}

func main() {
	// 创建一个拦截器
	interceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		fmt.Println("接收到了一个新的请求")

		// 从拦截器中获取metadata
		md, ok := metadata.FromIncomingContext(ctx)
		fmt.Println(md)
		if !ok {
			//已经开始接触到grpc的错误处理了
			return resp, status.Error(codes.Unauthenticated, "无token认证信息")
		}

		// 从拦截器中获取token认证信息
		var (
			appid  string
			appkey string
		)

		if va1, ok := md["appid"]; ok {
			appid = va1[0]
		}

		if va1, ok := md["appkey"]; ok {
			appkey = va1[0]
		}

		// 校验token认证信息
		if appid != "zhangdapeng" || appkey != "zhangdapeng520" {
			return resp, status.Error(codes.Unauthenticated, "无token认证信息")
		}

		// 执行请求
		res, err := handler(ctx, req)
		fmt.Println("请求已经完成")

		// 返回响应
		return res, err
	}

	// 使用拦截器
	opt := grpc.UnaryInterceptor(interceptor)

	// 创建服务
	g := grpc.NewServer(opt)

	// 注册服务
	proto.RegisterGreeterServer(g, &Server{})

	// 监听端口
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
