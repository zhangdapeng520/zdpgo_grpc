package main

/*
@Time : 2022/5/21 15:17
@Author : 张大鹏
@File : main.go
@Software: Goland2021.3.1
@Description:
*/

import (
	"basic/g"
	pb "basic/proto"
	"basic/server"
	"google.golang.org/grpc"
)

func registerFunc(s grpc.ServiceRegistrar, srv interface{}) {
	// 需要将传进来的对象转换为真实的GRPC服务对象
	pb.RegisterGreeterServer(s, srv.(pb.GreeterServer))
}

func main() {
	// 创建权限校验拦截器
	authInterceptor := g.Grpc.GetAuthInterceptor("zhangdapeng111", "zhangdapeng222")

	// 运行服务
	g.Grpc.RunServer(registerFunc, &server.Server{}, authInterceptor)
}
